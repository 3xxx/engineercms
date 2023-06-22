import { mat3, mat4 } from 'gl-matrix';

import * as macro from 'vtk.js/Sources/macros';
import vtkMapper from 'vtk.js/Sources/Rendering/Core/Mapper';
import vtkProperty from 'vtk.js/Sources/Rendering/Core/Property';
import vtkTexture from 'vtk.js/Sources/Rendering/Core/Texture';
import vtkWebGPUBufferManager from 'vtk.js/Sources/Rendering/WebGPU/BufferManager';
import vtkWebGPUShaderCache from 'vtk.js/Sources/Rendering/WebGPU/ShaderCache';
import vtkWebGPUUniformBuffer from 'vtk.js/Sources/Rendering/WebGPU/UniformBuffer';
import vtkWebGPUMapperHelper from 'vtk.js/Sources/Rendering/WebGPU/MapperHelper';
import vtkViewNode from 'vtk.js/Sources/Rendering/SceneGraph/ViewNode';

import { registerOverride } from 'vtk.js/Sources/Rendering/WebGPU/ViewNodeFactory';

const { BufferUsage, PrimitiveTypes } = vtkWebGPUBufferManager;
const { Representation } = vtkProperty;
const { ScalarMode } = vtkMapper;
const StartEvent = { type: 'StartEvent' };
const EndEvent = { type: 'EndEvent' };

const vtkWebGPUPolyDataVS = `
//VTK::Renderer::Dec

//VTK::Color::Dec

//VTK::Normal::Dec

//VTK::TCoord::Dec

//VTK::Select::Dec

//VTK::Mapper::Dec

//VTK::IOStructs::Dec

[[stage(vertex)]]
fn main(
//VTK::IOStructs::Input
)
//VTK::IOStructs::Output
{
  var output : vertexOutput;

  var vertex: vec4<f32> = vertexBC;

  //VTK::Color::Impl

  //VTK::Normal::Impl

  //VTK::TCoord::Impl

  //VTK::Select::Impl

  //VTK::Position::Impl

  return output;
}
`;

const vtkWebGPUPolyDataFS = `
//VTK::Renderer::Dec

//VTK::Color::Dec

// optional surface normal declaration
//VTK::Normal::Dec

//VTK::TCoord::Dec

//VTK::Select::Dec

//VTK::RenderEncoder::Dec

//VTK::Mapper::Dec

//VTK::IOStructs::Dec

[[stage(fragment)]]
fn main(
//VTK::IOStructs::Input
)
//VTK::IOStructs::Output
{
  var output : fragmentOutput;

  var ambientColor: vec4<f32> = mapperUBO.AmbientColor;
  var diffuseColor: vec4<f32> = mapperUBO.DiffuseColor;
  var opacity: f32 = mapperUBO.Opacity;

  //VTK::Color::Impl

  //VTK::Normal::Impl

  //VTK::Light::Impl

  var computedColor: vec4<f32> = vec4<f32>(ambientColor.rgb * mapperUBO.AmbientIntensity
     + diffuse * mapperUBO.DiffuseIntensity
     + specular * mapperUBO.SpecularIntensity,
     opacity);

  //VTK::TCoord::Impl

  //VTK::Select::Impl

  if (computedColor.a == 0.0) { discard; };

  //VTK::Position::Impl

  //VTK::RenderEncoder::Impl
  return output;
}
`;

function isEdges(hash) {
  // edge pipelines have "edge" in them
  return hash.indexOf('edge') >= 0;
}

// ----------------------------------------------------------------------------
// vtkWebGPUPolyDataMapper methods
// ----------------------------------------------------------------------------

function vtkWebGPUPolyDataMapper(publicAPI, model) {
  // Set our className
  model.classHierarchy.push('vtkWebGPUPolyDataMapper');

  publicAPI.buildPass = (prepass) => {
    if (prepass) {
      model.WebGPUActor = publicAPI.getFirstAncestorOfType('vtkWebGPUActor');
      model.WebGPURenderer =
        model.WebGPUActor.getFirstAncestorOfType('vtkWebGPURenderer');
      model.WebGPURenderWindow = model.WebGPURenderer.getParent();
      model.device = model.WebGPURenderWindow.getDevice();
    }
  };

  // Renders myself
  publicAPI.translucentPass = (prepass) => {
    if (prepass) {
      publicAPI.render();
    }
  };

  publicAPI.opaquePass = (prepass) => {
    if (prepass) {
      publicAPI.render();
    }
  };

  publicAPI.updateUBO = () => {
    // make sure the data is up to date
    const actor = model.WebGPUActor.getRenderable();
    const ppty = actor.getProperty();
    const utime = model.UBO.getSendTime();
    if (
      publicAPI.getMTime() > utime ||
      ppty.getMTime() > utime ||
      model.renderable.getMTime() > utime
    ) {
      const keyMats = model.WebGPUActor.getKeyMatrices(model.WebGPURenderer);
      model.UBO.setArray('BCWCMatrix', keyMats.bcwc);
      model.UBO.setArray('BCSCMatrix', keyMats.bcsc);
      model.UBO.setArray('MCWCNormals', keyMats.normalMatrix);

      let aColor = ppty.getAmbientColorByReference();
      model.UBO.setValue('AmbientIntensity', ppty.getAmbient());
      model.UBO.setArray('AmbientColor', [
        aColor[0],
        aColor[1],
        aColor[2],
        1.0,
      ]);
      model.UBO.setValue('DiffuseIntensity', ppty.getDiffuse());
      aColor = ppty.getDiffuseColorByReference();
      model.UBO.setArray('DiffuseColor', [
        aColor[0],
        aColor[1],
        aColor[2],
        1.0,
      ]);
      model.UBO.setValue('SpecularIntensity', ppty.getSpecular());
      model.UBO.setValue('SpecularPower', ppty.getSpecularPower());
      aColor = ppty.getSpecularColorByReference();
      model.UBO.setArray('SpecularColor', [
        aColor[0],
        aColor[1],
        aColor[2],
        1.0,
      ]);
      aColor = ppty.getEdgeColorByReference();
      model.UBO.setArray('EdgeColor', [aColor[0], aColor[1], aColor[2], 1.0]);
      model.UBO.setValue('Opacity', ppty.getOpacity());
      model.UBO.setValue('PropID', model.WebGPUActor.getPropID());
      const device = model.WebGPURenderWindow.getDevice();
      model.UBO.sendIfNeeded(device);
    }
  };

  publicAPI.render = () => {
    publicAPI.invokeEvent(StartEvent);
    if (!model.renderable.getStatic()) {
      model.renderable.update();
    }
    model.currentInput = model.renderable.getInputData();
    publicAPI.invokeEvent(EndEvent);

    model.renderEncoder = model.WebGPURenderer.getRenderEncoder();

    publicAPI.buildPrimitives();

    // update descriptor sets
    publicAPI.updateUBO();
  };

  publicAPI.replaceShaderPosition = (hash, pipeline, vertexInput) => {
    const vDesc = pipeline.getShaderDescription('vertex');
    vDesc.addBuiltinOutput('vec4<f32>', '[[builtin(position)]] Position');
    let code = vDesc.getCode();
    if (isEdges(hash)) {
      vDesc.addBuiltinInput('u32', '[[builtin(instance_index)]] instanceIndex');
      // widen the edge
      code = vtkWebGPUShaderCache.substitute(code, '//VTK::Position::Impl', [
        '    var tmpPos: vec4<f32> = rendererUBO.SCPCMatrix*mapperUBO.BCSCMatrix*vertexBC;',
        '    var tmpPos2: vec3<f32> = tmpPos.xyz / tmpPos.w;',
        '    tmpPos2.x = tmpPos2.x + 1.4*(f32(input.instanceIndex % 2u) - 0.5)/rendererUBO.viewportSize.x;',
        '    tmpPos2.y = tmpPos2.y + 1.4*(f32(input.instanceIndex / 2u) - 0.5)/rendererUBO.viewportSize.y;',
        '    tmpPos2.z = tmpPos2.z + 0.00001;', // could become a setting
        '    output.Position = vec4<f32>(tmpPos2.xyz * tmpPos.w, tmpPos.w);',
      ]).result;
    } else {
      code = vtkWebGPUShaderCache.substitute(code, '//VTK::Position::Impl', [
        '    output.Position = rendererUBO.SCPCMatrix*mapperUBO.BCSCMatrix*vertexBC;',
      ]).result;
    }

    vDesc.setCode(code);
  };

  publicAPI.replaceShaderNormal = (hash, pipeline, vertexInput) => {
    if (vertexInput.hasAttribute('normalMC')) {
      const vDesc = pipeline.getShaderDescription('vertex');
      vDesc.addOutput('vec3<f32>', 'normalVC');
      let code = vDesc.getCode();
      code = vtkWebGPUShaderCache.substitute(code, '//VTK::Normal::Impl', [
        '  output.normalVC = normalize((rendererUBO.WCVCNormals * mapperUBO.MCWCNormals * normalMC).xyz);',
      ]).result;
      vDesc.setCode(code);

      const fDesc = pipeline.getShaderDescription('fragment');
      code = fDesc.getCode();
      code = vtkWebGPUShaderCache.substitute(code, '//VTK::Normal::Impl', [
        '  var normal: vec3<f32> = input.normalVC;',
        '  if (!input.frontFacing) { normal = -normal; }',
      ]).result;
      fDesc.setCode(code);
    }
  };

  // we only apply lighting when there is a "var normal" declaration in the
  // fragment shader code. That is the lighting trigger.
  publicAPI.replaceShaderLight = (hash, pipeline, vertexInput) => {
    const fDesc = pipeline.getShaderDescription('fragment');
    let code = fDesc.getCode();
    if (code.includes('var normal')) {
      code = vtkWebGPUShaderCache.substitute(code, '//VTK::Light::Impl', [
        '  var df: f32  = max(0.0, normal.z);',
        '  var sf: f32 = pow(df, mapperUBO.SpecularPower);',
        '  var diffuse: vec3<f32> = df * diffuseColor.rgb;',
        '  var specular: vec3<f32> = sf * mapperUBO.SpecularColor.rgb * mapperUBO.SpecularColor.a;',
      ]).result;
      fDesc.setCode(code);
    } else {
      code = vtkWebGPUShaderCache.substitute(code, '//VTK::Light::Impl', [
        '  var diffuse: vec3<f32> = diffuseColor.rgb;',
        '  var specular: vec3<f32> = mapperUBO.SpecularColor.rgb * mapperUBO.SpecularColor.a;',
      ]).result;
      fDesc.setCode(code);
    }
  };

  publicAPI.replaceShaderColor = (hash, pipeline, vertexInput) => {
    if (isEdges(hash)) {
      const fDesc = pipeline.getShaderDescription('fragment');
      let code = fDesc.getCode();
      code = vtkWebGPUShaderCache.substitute(code, '//VTK::Color::Impl', [
        'ambientColor = mapperUBO.EdgeColor;',
        'diffuseColor = mapperUBO.EdgeColor;',
      ]).result;
      fDesc.setCode(code);
      return;
    }

    if (!vertexInput.hasAttribute('colorVI')) return;

    const vDesc = pipeline.getShaderDescription('vertex');
    vDesc.addOutput('vec4<f32>', 'color');
    let code = vDesc.getCode();
    code = vtkWebGPUShaderCache.substitute(code, '//VTK::Color::Impl', [
      '  output.color = colorVI;',
    ]).result;
    vDesc.setCode(code);

    const fDesc = pipeline.getShaderDescription('fragment');
    code = fDesc.getCode();
    code = vtkWebGPUShaderCache.substitute(code, '//VTK::Color::Impl', [
      'ambientColor = input.color;',
      'diffuseColor = input.color;',
      'opacity = mapperUBO.Opacity * input.color.a;',
    ]).result;
    fDesc.setCode(code);
  };

  publicAPI.replaceShaderTCoord = (hash, pipeline, vertexInput) => {
    if (!vertexInput.hasAttribute('tcoord')) return;

    const vDesc = pipeline.getShaderDescription('vertex');
    vDesc.addOutput('vec2<f32>', 'tcoordVS');
    let code = vDesc.getCode();
    code = vtkWebGPUShaderCache.substitute(code, '//VTK::TCoord::Impl', [
      '  output.tcoordVS = tcoord;',
    ]).result;
    vDesc.setCode(code);

    const fDesc = pipeline.getShaderDescription('fragment');
    code = fDesc.getCode();

    // todo handle multiple textures? Blend multiply ?
    if (model.textures.length) {
      code = vtkWebGPUShaderCache.substitute(code, '//VTK::TCoord::Impl', [
        'var tcolor: vec4<f32> = textureSample(Texture0, Texture0Sampler, input.tcoordVS);',
        'computedColor = computedColor*tcolor;',
      ]).result;
    }
    fDesc.setCode(code);
  };

  publicAPI.replaceShaderSelect = (hash, pipeline, vertexInput) => {
    if (hash.includes('sel')) {
      const fDesc = pipeline.getShaderDescription('fragment');
      let code = fDesc.getCode();
      // by default there are no composites, so just 0
      code = vtkWebGPUShaderCache.substitute(code, '//VTK::Select::Impl', [
        '  var compositeID: u32 = 0u;',
      ]).result;
      fDesc.setCode(code);
    }
  };

  publicAPI.getUsage = (rep, i) => {
    if (rep === Representation.POINTS || i === PrimitiveTypes.Points) {
      return BufferUsage.Verts;
    }

    if (i === PrimitiveTypes.Lines) {
      return BufferUsage.Lines;
    }

    if (rep === Representation.WIREFRAME) {
      if (i === PrimitiveTypes.Triangles) {
        return BufferUsage.LinesFromTriangles;
      }
      return BufferUsage.LinesFromStrips;
    }

    if (i === PrimitiveTypes.Triangles) {
      return BufferUsage.Triangles;
    }

    if (i === PrimitiveTypes.TriangleStrips) {
      return BufferUsage.Strips;
    }

    if (i === PrimitiveTypes.TriangleEdges) {
      return BufferUsage.LinesFromTriangles;
    }

    // only strip edges left which are lines
    return BufferUsage.LinesFromStrips;
  };

  publicAPI.getHashFromUsage = (usage) => `pt${usage}`;

  publicAPI.getTopologyFromUsage = (usage) => {
    switch (usage) {
      case BufferUsage.Triangles:
        return 'triangle-list';
      case BufferUsage.Verts:
        return 'point-list';
      default:
      case BufferUsage.Lines:
        return 'line-list';
    }
  };

  publicAPI.buildVertexInput = (pd, cells, primType) => {
    const actor = model.WebGPUActor.getRenderable();
    let representation = actor.getProperty().getRepresentation();
    const device = model.WebGPURenderWindow.getDevice();
    let edges = false;
    if (primType === PrimitiveTypes.TriangleEdges) {
      edges = true;
      representation = Representation.WIREFRAME;
    }

    const vertexInput = model.primitives[primType].getVertexInput();

    // hash = all things that can change the values on the buffer
    // since mtimes are unique we can use
    // - cells mtime - because cells drive how we pack
    // - rep (point/wireframe/surface) - again because of packing
    // - relevant dataArray mtime - the source data
    // - shift - not currently captured
    // - scale - not currently captured
    // - format
    // - usage
    // - packExtra - covered by format
    // - prim type (vert/lines/polys/strips) - covered by cells mtime

    const hash = cells.getMTime() + representation;
    // points
    const points = pd.getPoints();
    if (points) {
      const shift = model.WebGPUActor.getBufferShift(model.WebGPURenderer);
      const buffRequest = {
        hash: hash + points.getMTime(),
        dataArray: points,
        source: points,
        cells,
        primitiveType: primType,
        representation,
        time: Math.max(
          points.getMTime(),
          cells.getMTime(),
          model.WebGPUActor.getKeyMatricesTime().getMTime()
        ),
        shift,
        usage: BufferUsage.PointArray,
        format: 'float32x4',
        packExtra: true,
      };
      const buff = device.getBufferManager().getBuffer(buffRequest);
      vertexInput.addBuffer(buff, ['vertexBC']);
    } else {
      vertexInput.removeBufferIfPresent('vertexBC');
    }

    // normals, only used for surface rendering
    const usage = publicAPI.getUsage(representation, primType);
    if (usage === BufferUsage.Triangles || usage === BufferUsage.Strips) {
      const normals = pd.getPointData().getNormals();
      const buffRequest = {
        cells,
        representation,
        primitiveType: primType,
        format: 'snorm8x4',
        packExtra: true,
        shift: 0,
        scale: 127,
      };
      if (normals) {
        buffRequest.hash = hash + normals.getMTime();
        buffRequest.dataArray = normals;
        buffRequest.source = normals;
        buffRequest.time = Math.max(normals.getMTime(), cells.getMTime());
        buffRequest.usage = BufferUsage.PointArray;
        const buff = device.getBufferManager().getBuffer(buffRequest);
        vertexInput.addBuffer(buff, ['normalMC']);
      } else if (primType === PrimitiveTypes.Triangles) {
        buffRequest.hash = hash + points.getMTime();
        buffRequest.dataArray = points;
        buffRequest.source = points;
        buffRequest.time = Math.max(points.getMTime(), cells.getMTime());
        buffRequest.usage = BufferUsage.NormalsFromPoints;
        const buff = device.getBufferManager().getBuffer(buffRequest);
        vertexInput.addBuffer(buff, ['normalMC']);
      } else {
        vertexInput.removeBufferIfPresent('normalMC');
      }
    } else {
      vertexInput.removeBufferIfPresent('normalMC');
    }

    // deal with colors but only if modified
    let haveColors = false;
    if (model.renderable.getScalarVisibility()) {
      const c = model.renderable.getColorMapColors();
      if (c && !edges) {
        const scalarMode = model.renderable.getScalarMode();
        let haveCellScalars = false;
        // We must figure out how the scalars should be mapped to the polydata.
        if (
          (scalarMode === ScalarMode.USE_CELL_DATA ||
            scalarMode === ScalarMode.USE_CELL_FIELD_DATA ||
            scalarMode === ScalarMode.USE_FIELD_DATA ||
            !pd.getPointData().getScalars()) &&
          scalarMode !== ScalarMode.USE_POINT_FIELD_DATA &&
          c
        ) {
          haveCellScalars = true;
        }
        const buffRequest = {
          hash: hash + points.getMTime(),
          dataArray: c,
          source: c,
          cells,
          primitiveType: primType,
          representation,
          time: Math.max(c.getMTime(), cells.getMTime()),
          usage: BufferUsage.PointArray,
          format: 'unorm8x4',
          cellData: haveCellScalars,
          cellOffset: 0,
        };
        const buff = device.getBufferManager().getBuffer(buffRequest);
        vertexInput.addBuffer(buff, ['colorVI']);
        haveColors = true;
      }
    }
    if (!haveColors) {
      vertexInput.removeBufferIfPresent('colorVI');
    }

    let tcoords = null;
    if (
      model.renderable.getInterpolateScalarsBeforeMapping() &&
      model.renderable.getColorCoordinates()
    ) {
      tcoords = model.renderable.getColorCoordinates();
    } else {
      tcoords = pd.getPointData().getTCoords();
    }
    if (tcoords && !edges) {
      const buffRequest = {
        hash: hash + tcoords.getMTime(),
        dataArray: tcoords,
        source: tcoords,
        cells,
        primitiveType: primType,
        representation,
        time: Math.max(tcoords.getMTime(), cells.getMTime()),
        usage: BufferUsage.PointArray,
        format: 'float32x2',
      };
      const buff = device.getBufferManager().getBuffer(buffRequest);
      vertexInput.addBuffer(buff, ['tcoord']);
    } else {
      vertexInput.removeBufferIfPresent('tcoord');
    }
  };

  publicAPI.updateTextures = () => {
    // we keep track of new and used textures so
    // that we can clean up any unused textures so we don't hold onto them
    const usedTextures = [];
    const newTextures = [];

    // do we have a scalar color texture
    const idata = model.renderable.getColorTextureMap(); // returns an imagedata
    if (idata) {
      if (!model.colorTexture) {
        model.colorTexture = vtkTexture.newInstance();
      }
      model.colorTexture.setInputData(idata);
      newTextures.push(model.colorTexture);
    }

    // actor textures?
    const actor = model.WebGPUActor.getRenderable();
    const textures = actor.getTextures();
    for (let i = 0; i < textures.length; i++) {
      if (textures[i].getInputData()) {
        newTextures.push(textures[i]);
      }
      if (textures[i].getImage() && textures[i].getImageLoaded()) {
        newTextures.push(textures[i]);
      }
    }

    let usedCount = 0;
    for (let i = 0; i < newTextures.length; i++) {
      const srcTexture = newTextures[i];
      const treq = {};
      if (srcTexture.getInputData()) {
        treq.imageData = srcTexture.getInputData();
        treq.source = treq.imageData;
      } else if (srcTexture.getImage()) {
        treq.image = srcTexture.getImage();
        treq.source = treq.image;
      }
      const newTex = model.device.getTextureManager().getTexture(treq);
      if (newTex.getReady()) {
        // is this a new texture
        let found = false;
        for (let t = 0; t < model.textures.length; t++) {
          if (model.textures[t] === newTex) {
            usedCount++;
            found = true;
            usedTextures[t] = true;
          }
        }
        if (!found) {
          usedTextures[model.textures.length] = true;
          const tview = newTex.createView();
          tview.setName(`Texture${usedCount++}`);
          model.textures.push(newTex);
          model.textureViews.push(tview);
          const interpolate = srcTexture.getInterpolate()
            ? 'linear'
            : 'nearest';
          tview.addSampler(model.device, {
            minFilter: interpolate,
            magFilter: interpolate,
          });
        }
      }
    }

    // remove unused textures
    for (let i = model.textures.length - 1; i >= 0; i--) {
      if (!usedTextures[i]) {
        model.textures.splice(i, 1);
        model.textureViews.splice(i, 1);
      }
    }
  };

  // compute a unique hash for a pipeline, this needs to be unique enough to
  // capture any pipeline code changes (which includes shader changes)
  // or vertex input changes/ bind groups/ etc
  publicAPI.computePipelineHash = (vertexInput, usage, edges) => {
    let pipelineHash = 'pd';
    if (edges) {
      pipelineHash += 'edge';
    } else {
      if (vertexInput.hasAttribute(`normalMC`)) {
        pipelineHash += `n`;
      }
      if (vertexInput.hasAttribute(`colorVI`)) {
        pipelineHash += `c`;
      }
      if (vertexInput.hasAttribute(`tcoord`)) {
        pipelineHash += `t`;
      }
      if (model.textures.length) {
        pipelineHash += `tx${model.textures.length}`;
      }
    }

    if (model.SSBO) {
      pipelineHash += `ssbo`;
    }
    const uhash = publicAPI.getHashFromUsage(usage);
    pipelineHash += uhash;
    pipelineHash += model.renderEncoder.getPipelineHash();

    return pipelineHash;
  };

  // was originally buildIBOs() but not using IBOs right now
  publicAPI.buildPrimitives = () => {
    const poly = model.currentInput;
    const prims = [
      poly.getVerts(),
      poly.getLines(),
      poly.getPolys(),
      poly.getStrips(),
    ];

    const device = model.WebGPURenderWindow.getDevice();

    model.renderable.mapScalars(poly, 1.0);

    // handle textures
    publicAPI.updateTextures();

    const actor = model.WebGPUActor.getRenderable();
    const rep = actor.getProperty().getRepresentation();
    const edgeVisibility = actor.getProperty().getEdgeVisibility();

    // handle per primitive type
    for (let i = PrimitiveTypes.Points; i <= PrimitiveTypes.Triangles; i++) {
      if (prims[i].getNumberOfValues() > 0) {
        {
          const usage = publicAPI.getUsage(rep, i);
          const primHelper = model.primitives[i];

          publicAPI.buildVertexInput(model.currentInput, prims[i], i);
          primHelper.setPipelineHash(
            publicAPI.computePipelineHash(
              primHelper.getVertexInput(),
              usage,
              false
            )
          );

          primHelper.setTextureViews(model.textureViews);
          primHelper.setWebGPURenderer(model.WebGPURenderer);
          primHelper.setNumberOfInstances(1);
          const vbo = primHelper.getVertexInput().getBuffer('vertexBC');
          primHelper.setNumberOfVertices(
            vbo.getSizeInBytes() / vbo.getStrideInBytes()
          );
          primHelper.setTopology(publicAPI.getTopologyFromUsage(usage));
          primHelper.build(model.renderEncoder, device);
          primHelper.registerToDraw();
        }

        // also handle edge visibility if turned on
        if (
          edgeVisibility &&
          rep === Representation.SURFACE &&
          i === PrimitiveTypes.Triangles
        ) {
          const primHelper = model.primitives[PrimitiveTypes.TriangleEdges];
          const usage = publicAPI.getUsage(rep, PrimitiveTypes.TriangleEdges);

          publicAPI.buildVertexInput(
            model.currentInput,
            prims[PrimitiveTypes.Triangles],
            PrimitiveTypes.TriangleEdges
          );
          primHelper.setPipelineHash(
            publicAPI.computePipelineHash(
              primHelper.getVertexInput(),
              usage,
              true
            )
          );

          primHelper.setWebGPURenderer(model.WebGPURenderer);
          primHelper.setNumberOfInstances(4);
          const vbo = primHelper.getVertexInput().getBuffer('vertexBC');
          primHelper.setNumberOfVertices(
            vbo.getSizeInBytes() / vbo.getStrideInBytes()
          );
          primHelper.setTopology(publicAPI.getTopologyFromUsage(usage));
          primHelper.build(model.renderEncoder, device);
          primHelper.registerToDraw();
        }
      }
    }
  };

  publicAPI.setShaderReplacement = (name, func) => {
    for (let i = PrimitiveTypes.Start; i < PrimitiveTypes.End; i++) {
      const sr = model.primitives[i].getShaderReplacements();
      sr.set(name, func);
    }
  };

  publicAPI.setFragmentShaderTemplate = (val) => {
    model.fragmentShaderTemplate = val;
    for (let i = PrimitiveTypes.Start; i < PrimitiveTypes.End; i++) {
      model.primitives[i].setFragmentShaderTemplate(val);
    }
  };

  publicAPI.setVertexShaderTemplate = (val) => {
    model.fragmentShaderTemplate = val;
    for (let i = PrimitiveTypes.Start; i < PrimitiveTypes.End; i++) {
      model.primitives[i].setVertexShaderTemplate(val);
    }
  };
}

// ----------------------------------------------------------------------------
// Object factory
// ----------------------------------------------------------------------------

const DEFAULT_VALUES = {
  colorTexture: null,
  renderEncoder: null,
  textures: null,
  textureViews: null,
  primitives: null,
  tmpMat4: null,
  fragmentShaderTemplate: null,
  vertexShaderTemplate: null,
};

// ----------------------------------------------------------------------------

export function extend(publicAPI, model, initialValues = {}) {
  Object.assign(model, DEFAULT_VALUES, initialValues);

  // Inheritance
  vtkViewNode.extend(publicAPI, model, initialValues);

  model.tmpMat3 = mat3.identity(new Float64Array(9));
  model.tmpMat4 = mat4.identity(new Float64Array(16));

  model.fragmentShaderTemplate =
    model.fragmentShaderTemplate || vtkWebGPUPolyDataFS;
  model.vertexShaderTemplate =
    model.vertexShaderTemplate || vtkWebGPUPolyDataVS;

  model.UBO = vtkWebGPUUniformBuffer.newInstance();
  model.UBO.setName('mapperUBO');
  model.UBO.addEntry('BCWCMatrix', 'mat4x4<f32>');
  model.UBO.addEntry('BCSCMatrix', 'mat4x4<f32>');
  model.UBO.addEntry('MCWCNormals', 'mat4x4<f32>');
  model.UBO.addEntry('AmbientColor', 'vec4<f32>');
  model.UBO.addEntry('DiffuseColor', 'vec4<f32>');
  model.UBO.addEntry('EdgeColor', 'vec4<f32>');
  model.UBO.addEntry('AmbientIntensity', 'f32');
  model.UBO.addEntry('DiffuseIntensity', 'f32');
  model.UBO.addEntry('SpecularColor', 'vec4<f32>');
  model.UBO.addEntry('SpecularIntensity', 'f32');
  model.UBO.addEntry('Opacity', 'f32');
  model.UBO.addEntry('SpecularPower', 'f32');
  model.UBO.addEntry('PropID', 'u32');

  // Build VTK API
  macro.get(publicAPI, model, [
    'fragmentShaderTemplate',
    'vertexShaderTemplate',
    'UBO',
  ]);
  macro.setGet(publicAPI, model, ['renderEncoder']);

  model.textures = [];
  model.textureViews = [];
  model.primitives = [];

  // Object methods
  vtkWebGPUPolyDataMapper(publicAPI, model);

  for (let i = PrimitiveTypes.Start; i < PrimitiveTypes.End; i++) {
    model.primitives[i] = vtkWebGPUMapperHelper.newInstance();
    model.primitives[i].setUBO(model.UBO);
    model.primitives[i].setVertexShaderTemplate(
      publicAPI.getVertexShaderTemplate()
    );
    model.primitives[i].setFragmentShaderTemplate(
      publicAPI.getFragmentShaderTemplate()
    );
  }

  publicAPI.setShaderReplacement(
    'replaceShaderPosition',
    publicAPI.replaceShaderPosition
  );
  publicAPI.setShaderReplacement(
    'replaceShaderLight',
    publicAPI.replaceShaderLight
  );
  publicAPI.setShaderReplacement(
    'replaceShaderTCoord',
    publicAPI.replaceShaderTCoord
  );
  publicAPI.setShaderReplacement(
    'replaceShaderNormal',
    publicAPI.replaceShaderNormal
  );
  publicAPI.setShaderReplacement(
    'replaceShaderSelect',
    publicAPI.replaceShaderSelect
  );
  publicAPI.setShaderReplacement(
    'replaceShaderColor',
    publicAPI.replaceShaderColor
  );
}

// ----------------------------------------------------------------------------

export const newInstance = macro.newInstance(extend, 'vtkWebGPUPolyDataMapper');

// ----------------------------------------------------------------------------

export default { newInstance, extend };

// Register ourself to WebGPU backend if imported
registerOverride('vtkMapper', newInstance);
