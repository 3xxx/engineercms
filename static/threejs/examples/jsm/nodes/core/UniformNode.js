import InputNode from './InputNode.js';

class UniformNode extends InputNode {

	getUniformHash( builder ) {

		return this.getHash( builder );

	}

	generate( builder, output ) {

		const type = this.getNodeType( builder );

		const hash = this.getUniformHash( builder );

		let sharedNode = builder.getNodeFromHash( hash );

		if ( sharedNode === undefined ) {

			builder.setHashNode( this, hash );

			sharedNode = this;

		}

		const sharedNodeType = sharedNode.getInputType( builder );

		const nodeUniform = builder.getUniformFromNode( sharedNode, builder.shaderStage, sharedNodeType );
		const propertyName = builder.getPropertyName( nodeUniform );

		return builder.format( propertyName, type, output );

	}

}

UniformNode.prototype.isUniformNode = true;

export default UniformNode;
