import Node from './Node.js';

class BypassNode extends Node {

	constructor( returnNode, callNode ) {

		super();

		this.outputNode = returnNode;
		this.callNode = callNode;

	}

	getNodeType( builder ) {

		return this.outputNode.getNodeType( builder );

	}

	generate( builder, output ) {

		const snippet = this.callNode.build( builder, 'void' );

		if ( snippet !== '' ) {

			builder.addFlowCode( snippet );

		}

		return this.outputNode.build( builder, output );

	}

}

BypassNode.prototype.isBypassNode = true;

export default BypassNode;
