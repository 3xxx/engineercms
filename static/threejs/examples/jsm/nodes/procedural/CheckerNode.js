import Node from '../core/Node.js';

import { ShaderNode, uv, add, mul, floor, mod, sign } from '../ShaderNode.js';

const checkerShaderNode = new ShaderNode( ( inputs ) => {

	const uv = mul( inputs.uv, 2.0 );

	const cx = floor( uv.x );
	const cy = floor( uv.y );
	const result = mod( add( cx, cy ), 2.0 );

	return sign( result );

} );

class CheckerNode extends Node {

	constructor( uvNode = uv() ) {

		super( 'float' );

		this.uvNode = uvNode;

	}

	generate( builder ) {

		return checkerShaderNode( { uv: this.uvNode } ).build( builder );

	}

}

export default CheckerNode;
