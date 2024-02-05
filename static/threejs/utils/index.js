import * as THREE from 'three'

import { RenderPass } from '/static/threejs/examples/jsm/postprocessing/RenderPass.js';
import { EffectComposer } from '/static/threejs/examples/jsm/postprocessing/EffectComposer.js';
import { UnrealBloomPass } from '/static/threejs/examples/jsm/postprocessing/UnrealBloomPass.js';
import { OutputPass } from '/static/threejs/examples/jsm/postprocessing/OutputPass.js';
import { ShaderPass } from '/static/threejs/examples/jsm/postprocessing/ShaderPass.js';


let color = new THREE.Color('#0fb1fb')

const meshMaterial = new THREE.MeshBasicMaterial({
  color: color,
  transparent: true,
  opacity: 0.3
});
export const changeModelMaterial = (object, lineGroup) => {
  var object = THREE.Object3D
  var lineGroup = THREE.Group
  var group = THREE.Group
  // 因为和THREE.Object3D并无差异，所以var group = new THREE.Group();可以使用var group = new THREE.Object3D(); 
  group = object
  if (group.isObject3D) {
    const lg = new THREE.Group()
    lineGroup.add(lg)
    lg.name = group.name + '_line'
    group.traverse((mesh) => {
      var mesh = Mesh
      if (mesh.isMesh) {
        const quaternion = new THREE.Quaternion()
        const worldPos = new THREE.Vector3()
        const worldScale = new THREE.Vector3()
        // 获取四元数
        mesh.getWorldQuaternion(quaternion)
        // 获取位置信息
        mesh.getWorldPosition(worldPos)
        // 获取缩放比例
        mesh.getWorldScale(worldScale)

        mesh.material = meshMaterial

        // 以模型顶点信息创建线条
        const line = getLine(mesh, Math.PI * 6.137, undefined, 1)
        const name = mesh.name + '_line'

        line.name = name

        // 给线段赋予模型相同的坐标信息
        line.quaternion.copy(quaternion)
        line.position.copy(worldPos)
        line.scale.copy(worldScale)

        lg.add(line)
      }
    })

  }
}


const lineMaterial = new THREE.LineBasicMaterial({
  color: new THREE.Color(color),
  depthTest: true,
  transparent: true
})
export const getLine = (object, thresholdAngle = 1, color = new THREE.Color('#ff0ff0'), opacity = 1) => {
  var object = THREE.Mesh
  var getLine = THREE.LineSegments
  // 创建线条，参数为 几何体模型，相邻面的法线之间的角度，
  var edges = new THREE.EdgesGeometry(object.geometry, thresholdAngle);
  var line = new THREE.LineSegments(edges);
  lineMaterial.opacity = opacity
  line.material = lineMaterial
  return line;
}


const params = {
  threshold: 0,
  strength: 0.972, // 强度
  radius: 0.21, // 半径
  exposure: 1.55 // 扩散
};
export const unreal = (scene, camera, renderer, width, height) => {
  var scene = THREE.Scene
  var camera = THREE.PerspectiveCamera
  // var renderer= THREE.WebGLRenderer
  // Renderer
  var renderer = new THREE.WebGLRenderer({
    // canvas,
    antialias: true,
  })
  // : number: number
  // 渲染器通道，将场景全部加入渲染器
  const renderScene = new RenderPass(scene, camera);
  // 添加虚幻发光通道
  const bloomPass = new UnrealBloomPass(new THREE.Vector2(width, height), 1.5, 0.4, 0.85);
  bloomPass.threshold = params.threshold;
  bloomPass.strength = params.strength;
  bloomPass.radius = params.radius;

  // 创建合成器
  const bloomComposer = new EffectComposer(renderer);
  bloomComposer.renderToScreen = false;
  // 将渲染器和场景结合到合成器中
  bloomComposer.addPass(renderScene);
  bloomComposer.addPass(bloomPass);

  // 着色器通道
  const mixPass = new ShaderPass(
    // 着色器
    new THREE.ShaderMaterial({
      uniforms: {
        baseTexture: { value: null },
        bloomTexture: { value: bloomComposer.renderTarget2.texture }
      },
      vertexShader: `
            
            varying vec2 vUv;
    
            void main() {
    
                vUv = uv;
    
                gl_Position = projectionMatrix * modelViewMatrix * vec4( position, 1.0 );
    
            }
            
            `,
      fragmentShader: `
            
            uniform sampler2D baseTexture;
            uniform sampler2D bloomTexture;
    
            varying vec2 vUv;
    
            void main() {
    
                gl_FragColor = ( texture2D( baseTexture, vUv ) + vec4( 1.0 ) * texture2D( bloomTexture, vUv ) );
    
            }
    
            `,
      defines: {}
    }), 'baseTexture'
  );
  mixPass.needsSwap = true;

  // 合成器输出通道
  const outputPass = new OutputPass();

  const finalComposer = new EffectComposer(renderer);
  finalComposer.addPass(renderScene);
  finalComposer.addPass(mixPass);
  finalComposer.addPass(outputPass);

  return {
    finalComposer,
    // bloomComposer,
    renderScene,
    bloomPass
  }
}