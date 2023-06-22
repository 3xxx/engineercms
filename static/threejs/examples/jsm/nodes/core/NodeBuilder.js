import NodeUniform from './NodeUniform.js';
import NodeAttribute from './NodeAttribute.js';
import NodeVary from './NodeVary.js';
import NodeVar from './NodeVar.js';
import NodeCode from './NodeCode.js';
import NodeKeywords from './NodeKeywords.js';
import { NodeUpdateType } from './constants.js';

import { REVISION, LinearEncoding } from 'three';

export const shaderStages = [ 'fragment', 'vertex' ];
export const vector = [ 'x', 'y', 'z', 'w' ];

const toFloat = ( value ) => {
	
	value = Number( value );

	return value + ( value % 1 ? '' : '.0' );

};

class NodeBuilder {

	constructor( object, renderer, parser ) {

		this.object = object;
		this.material = object.material;
		this.renderer = renderer;
		this.parser = parser;

		this.nodes = [];
		this.updateNodes = [];
		this.hashNodes = {};

		this.vertexShader = null;
		this.fragmentShader = null;

		this.flowNodes = { vertex: [], fragment: [] };
		this.flowCode = { vertex: '', fragment: '' };
		this.uniforms = { vertex: [], fragment: [], index: 0 };
		this.codes = { vertex: [], fragment: [] };
		this.attributes = [];
		this.varys = [];
		this.vars = { vertex: [], fragment: [] };
		this.flow = { code: '' };
		this.stack = [];

		this.context = {
			keywords: new NodeKeywords(),
			material: object.material
		};

		this.nodesData = new WeakMap();
		this.flowsData = new WeakMap();

		this.shaderStage = null;
		this.node = null;

	}

	addStack( node ) {

		/*
		if ( this.stack.indexOf( node ) !== - 1 ) {

			console.warn( 'Recursive node: ', node );

		}
		*/

		this.stack.push( node );

	}

	removeStack( node ) {

		const lastStack = this.stack.pop();

		if ( lastStack !== node ) {

			throw new Error( 'NodeBuilder: Invalid node stack!' );

		}

	}

	setHashNode( node, hash ) {

		this.hashNodes[ hash ] = node;

	}

	addNode( node ) {

		if ( this.nodes.indexOf( node ) === - 1 ) {

			const updateType = node.getUpdateType( this );

			if ( updateType !== NodeUpdateType.None ) {

				this.updateNodes.push( node );

			}

			this.nodes.push( node );

			this.setHashNode( node, node.getHash( this ) );

		}

	}

	getMethod( method ) {

		return method;

	}

	getNodeFromHash( hash ) {

		return this.hashNodes[ hash ];

	}

	addFlow( shaderStage, node ) {

		this.flowNodes[ shaderStage ].push( node );

		return node;

	}

	setContext( context ) {

		this.context = context;

	}

	getContext() {

		return this.context;

	}

	getTexture( /* textureProperty, uvSnippet, biasSnippet = null */ ) {

		console.warn( 'Abstract function.' );

	}

	getCubeTexture( /* textureProperty, uvSnippet, biasSnippet = null */ ) {

		console.warn( 'Abstract function.' );

	}

	// @TODO: rename to .generateConst()
	getConst( type, value ) {

		if ( type === 'float' ) return toFloat( value );
		if ( type === 'int' ) return `${ Math.round( value ) }`;
		if ( type === 'uint' ) return value >= 0 ? `${ Math.round( value ) }` : '0';
		if ( type === 'bool' ) return value ? 'true' : 'false';
		if ( type === 'color' ) return `${ this.getType( 'vec3' ) }( ${ toFloat( value.r ) }, ${ toFloat( value.g ) }, ${ toFloat( value.b ) } )`;

		const typeLength = this.getTypeLength( type );

		const componentType = this.getComponentType( type );

		const getConst = value => this.getConst( componentType, value );

		if ( typeLength === 2 ) {

			return `${ this.getType( type ) }( ${ getConst( value.x ) }, ${ getConst( value.y ) } )`;

		} else if ( typeLength === 3 ) {

			return `${ this.getType( type ) }( ${ getConst( value.x ) }, ${ getConst( value.y ) }, ${ getConst( value.z ) } )`;

		} else if ( typeLength === 4 ) {

			return `${ this.getType( type ) }( ${ getConst( value.x ) }, ${ getConst( value.y ) }, ${ getConst( value.z ) }, ${ getConst( value.w ) } )`;

		}

		throw new Error( `NodeBuilder: Type '${type}' not found in generate constant attempt.` );

	}

	getType( type ) {

		return type;

	}

	generateMethod( method ) {

		return method;

	}

	getAttribute( name, type ) {

		const attributes = this.attributes;

		// find attribute

		for ( const attribute of attributes ) {

			if ( attribute.name === name ) {

				return attribute;

			}

		}

		// create a new if no exist

		const attribute = new NodeAttribute( name, type );

		attributes.push( attribute );

		return attribute;

	}

	getPropertyName( node/*, shaderStage*/ ) {

		return node.name;

	}

	isVector( type ) {

		return /vec\d/.test( type );

	}

	isMatrix( type ) {

		return /mat\d/.test( type );

	}

	isShaderStage( shaderStage ) {

		return this.shaderStage === shaderStage;

	}

	getTextureEncodingFromMap( map ) {

		let encoding;

		if ( map && map.isTexture ) {

			encoding = map.encoding;

		} else if ( map && map.isWebGLRenderTarget ) {

			encoding = map.texture.encoding;

		} else {

			encoding = LinearEncoding;

		}

		return encoding;

	}

	getComponentType( type ) {

		type = this.getVectorType( type );

		const componentType = /(b|i|u|)(vec|mat)([2-4])/.exec( type );

		if ( componentType === null ) return null;

		if ( componentType[ 1 ] === 'b' ) return 'bool';
		if ( componentType[ 1 ] === 'i' ) return 'int';
		if ( componentType[ 1 ] === 'u' ) return 'uint';

		return 'float';

	}

	getVectorType( type ) {

		if ( type === 'color' ) return 'vec3';
		if ( type === 'texture' ) return 'vec4';

		return type;

	}

	getTypeFromLength( type ) {

		if ( type === 1 ) return 'float';
		if ( type === 2 ) return 'vec2';
		if ( type === 3 ) return 'vec3';
		if ( type === 4 ) return 'vec4';

		return 0;

	}

	getTypeLength( type ) {

		const vecType = this.getVectorType( type );
		const vecNum = /vec([2-4])/.exec( vecType );

		if ( vecNum !== null ) return Number( vecNum[ 1 ] );
		if ( vecType === 'float' || vecType === 'bool' || vecType === 'int' || vecType === 'uint' ) return 1;

		return 0;

	}

	getVectorFromMatrix( type ) {

		return type.replace( 'mat', 'vec' );

	}

	getDataFromNode( node, shaderStage = this.shaderStage ) {

		let nodeData = this.nodesData.get( node );

		if ( nodeData === undefined ) {

			nodeData = { vertex: {}, fragment: {} };

			this.nodesData.set( node, nodeData );

		}

		return shaderStage !== null ? nodeData[ shaderStage ] : nodeData;

	}

	getUniformFromNode( node, shaderStage, type ) {

		const nodeData = this.getDataFromNode( node, shaderStage );

		let nodeUniform = nodeData.uniform;

		if ( nodeUniform === undefined ) {

			const index = this.uniforms.index ++;

			nodeUniform = new NodeUniform( 'nodeUniform' + index, type, node );

			this.uniforms[ shaderStage ].push( nodeUniform );

			nodeData.uniform = nodeUniform;

		}

		return nodeUniform;

	}

	getVarFromNode( node, type, shaderStage = this.shaderStage ) {

		const nodeData = this.getDataFromNode( node, shaderStage );

		let nodeVar = nodeData.variable;

		if ( nodeVar === undefined ) {

			const vars = this.vars[ shaderStage ];
			const index = vars.length;

			nodeVar = new NodeVar( 'nodeVar' + index, type );

			vars.push( nodeVar );

			nodeData.variable = nodeVar;

		}

		return nodeVar;

	}

	getVaryFromNode( node, type ) {

		const nodeData = this.getDataFromNode( node, null );

		let nodeVary = nodeData.vary;

		if ( nodeVary === undefined ) {

			const varys = this.varys;
			const index = varys.length;

			nodeVary = new NodeVary( 'nodeVary' + index, type );

			varys.push( nodeVary );

			nodeData.vary = nodeVary;

		}

		return nodeVary;

	}

	getCodeFromNode( node, type, shaderStage = this.shaderStage ) {

		const nodeData = this.getDataFromNode( node );

		let nodeCode = nodeData.code;

		if ( nodeCode === undefined ) {

			const codes = this.codes[ shaderStage ];
			const index = codes.length;

			nodeCode = new NodeCode( 'nodeCode' + index, type );

			codes.push( nodeCode );

			nodeData.code = nodeCode;

		}

		return nodeCode;

	}

	addFlowCode( code ) {

		this.flow.code += code;

	}

	getFlowData( shaderStage, node ) {

		return this.flowsData.get( node );

	}

	flowNode( node ) {

		this.node = node;

		const output = node.getNodeType( this );

		const flowData = this.flowChildNode( node, output );

		this.flowsData.set( node, flowData );

		this.node = null;

		return flowData;

	}

	flowChildNode( node, output = null ) {

		const previousFlow = this.flow;

		const flow = {
			code: '',
		};

		this.flow = flow;

		flow.result = node.build( this, output );

		this.flow = previousFlow;

		return flow;

	}

	flowNodeFromShaderStage( shaderStage, node, output = null, propertyName = null ) {

		const previousShaderStage = this.shaderStage;

		this.setShaderStage( shaderStage );

		const flowData = this.flowChildNode( node, output );

		if ( propertyName !== null ) {

			flowData.code += `${propertyName} = ${flowData.result};\n\t`;

		}

		this.flowCode[ shaderStage ] = this.flowCode[ shaderStage ] + flowData.code;

		this.setShaderStage( previousShaderStage );

		return flowData;

	}

	getAttributes( /*shaderStage*/ ) {

		console.warn( 'Abstract function.' );

	}

	getVarys( /*shaderStage*/ ) {

		console.warn( 'Abstract function.' );

	}

	getVars( shaderStage ) {

		let snippet = '';

		const vars = this.vars[ shaderStage ];

		for ( let index = 0; index < vars.length; index ++ ) {

			const variable = vars[ index ];

			snippet += `${variable.type} ${variable.name}; `;

		}

		return snippet;

	}

	getUniforms( /*shaderStage*/ ) {

		console.warn( 'Abstract function.' );

	}

	getCodes( shaderStage ) {

		const codes = this.codes[ shaderStage ];

		let code = '';

		for ( const nodeCode of codes ) {

			code += nodeCode.code + '\n';

		}

		return code;

	}

	getHash() {

		return this.vertexShader + this.fragmentShader;

	}

	getShaderStage() {

		return this.shaderStage;

	}

	setShaderStage( shaderStage ) {

		this.shaderStage = shaderStage;

	}

	buildCode() {

		console.warn( 'Abstract function.' );

	}

	build() {

		// stage 1: analyze nodes to possible optimization and validation

		for ( const shaderStage of shaderStages ) {

			this.setShaderStage( shaderStage );

			const flowNodes = this.flowNodes[ shaderStage ];

			for ( const node of flowNodes ) {

				node.analyze( this );

			}

		}

		// stage 2: pre-build vertex code used in fragment shader

		if ( this.context.vertex && this.context.vertex.isNode ) {

			this.flowNodeFromShaderStage( 'vertex', this.context.vertex );

		}

		// stage 3: generate shader

		for ( const shaderStage of shaderStages ) {

			this.setShaderStage( shaderStage );

			const flowNodes = this.flowNodes[ shaderStage ];

			for ( const node of flowNodes ) {

				this.flowNode( node, shaderStage );

			}

		}

		this.setShaderStage( null );

		// stage 4: build code for a specific output

		this.buildCode();

		return this;

	}

	format( snippet, fromType, toType ) {

		fromType = this.getVectorType( fromType );
		toType = this.getVectorType( toType );

		if ( fromType === toType || toType === 'void' || toType === null ) {

			return snippet;

		}

		const fromTypeLength = this.getTypeLength( fromType );
		const toTypeLength = this.getTypeLength( toType );

		if ( fromTypeLength === 0 ) { // fromType is matrix-like

			const vectorType = this.getVectorFromMatrix( fromType );

			return this.format( `( ${ snippet } * ${ this.getType( vectorType ) }( 1.0 ) )`, vectorType, toType );

		}

		if ( toTypeLength === 0 ) { // toType is matrix-like

			// ignore for now
			//return `${ this.getType( toType ) }( ${ snippet } )`;

			return snippet;

		}

		if ( fromTypeLength === toTypeLength ) {

			return `${ this.getType( toType ) }( ${ snippet } )`;

		}

		if ( fromTypeLength > toTypeLength ) {

			return this.format( `${ snippet }.${ 'xyz'.slice( 0, toTypeLength ) }`, this.getTypeFromLength( toTypeLength ), toType );

		}

		if ( toTypeLength === 4 ) { // toType is vec4-like

			return `${ this.getType( toType ) }( ${ this.format( snippet, fromType, 'vec3' ) }, 1.0 )`;

		}

		if ( fromTypeLength === 2 ) { // fromType is vec2-like and toType is vec3-like

			return `${ this.getType( toType ) }( ${ this.format( snippet, fromType, 'vec2' ) }, 0.0 )`;

		}

		return `${ this.getType( toType ) }( ${ snippet } )`; // fromType is float-like

	}

	getSignature() {

		return `// Three.js r${ REVISION } - NodeMaterial System\n`;

	}

}

export default NodeBuilder;
