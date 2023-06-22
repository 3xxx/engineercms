import { Command } from '../Command.js';

/**
 * @param editor Editor
 * @param object THREE.Object3D
 * @param newParent THREE.Object3D
 * @param newBefore THREE.Object3D
 * @constructor
 */
class MoveObjectCommand extends Command {

	constructor( editor, object, newParent, newBefore ) {

		super( editor );

		this.type = 'MoveObjectCommand';
		this.name = 'Move Object';

		this.object = object;
		this.oldParent = ( object !== undefined ) ? object.parent : undefined;
		this.oldIndex = ( this.oldParent !== undefined ) ? this.oldParent.children.indexOf( this.object ) : undefined;
		this.newParent = newParent;

		if ( newBefore !== undefined ) {

			this.newIndex = ( newParent !== undefined ) ? newParent.children.indexOf( newBefore ) : undefined;

		} else {

			this.newIndex = ( newParent !== undefined ) ? newParent.children.length : undefined;

		}

		if ( this.oldParent === this.newParent && this.newIndex > this.oldIndex ) {

			this.newIndex --;

		}

		this.newBefore = newBefore;

	}

	execute() {

		this.oldParent.remove( this.object );

		const children = this.newParent.children;
		children.splice( this.newIndex, 0, this.object );
		this.object.parent = this.newParent;

		this.object.dispatchEvent( { type: 'added' } );
		this.editor.signals.sceneGraphChanged.dispatch();

	}

	undo() {

		this.newParent.remove( this.object );

		const children = this.oldParent.children;
		children.splice( this.oldIndex, 0, this.object );
		this.object.parent = this.oldParent;

		this.object.dispatchEvent( { type: 'added' } );
		this.editor.signals.sceneGraphChanged.dispatch();

	}

	toJSON() {

		const output = super.toJSON( this );

		output.objectUuid = this.object.uuid;
		output.newParentUuid = this.newParent.uuid;
		output.oldParentUuid = this.oldParent.uuid;
		output.newIndex = this.newIndex;
		output.oldIndex = this.oldIndex;

		return output;

	}

	fromJSON( json ) {

		super.fromJSON( json );

		this.object = this.editor.objectByUuid( json.objectUuid );
		this.oldParent = this.editor.objectByUuid( json.oldParentUuid );
		if ( this.oldParent === undefined ) {

			this.oldParent = this.editor.scene;

		}

		this.newParent = this.editor.objectByUuid( json.newParentUuid );

		if ( this.newParent === undefined ) {

			this.newParent = this.editor.scene;

		}

		this.newIndex = json.newIndex;
		this.oldIndex = json.oldIndex;

	}

}

export { MoveObjectCommand };
