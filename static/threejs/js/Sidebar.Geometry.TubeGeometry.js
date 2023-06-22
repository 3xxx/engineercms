import * as THREE from 'three';

import { UIDiv, UIRow, UIText, UIInteger, UISelect, UICheckbox, UINumber } from './libs/ui.js';
import { UIPoints3 } from './libs/ui.three.js';

import { SetGeometryCommand } from './commands/SetGeometryCommand.js';

function GeometryParametersPanel( editor, object ) {

	const strings = editor.strings;

	const container = new UIDiv();

	const geometry = object.geometry;
	const parameters = geometry.parameters;

	// points

	const pointsRow = new UIRow();
	pointsRow.add( new UIText( strings.getKey( 'sidebar/geometry/tube_geometry/path' ) ).setWidth( '90px' ) );

	const points = new UIPoints3().setValue( parameters.path.points ).onChange( update );
	pointsRow.add( points );

	container.add( pointsRow );

	// radius

	const radiusRow = new UIRow();
	const radius = new UINumber( parameters.radius ).onChange( update );

	radiusRow.add( new UIText( strings.getKey( 'sidebar/geometry/tube_geometry/radius' ) ).setWidth( '90px' ) );
	radiusRow.add( radius );

	container.add( radiusRow );

	// tubularSegments

	const tubularSegmentsRow = new UIRow();
	const tubularSegments = new UIInteger( parameters.tubularSegments ).onChange( update );

	tubularSegmentsRow.add( new UIText( strings.getKey( 'sidebar/geometry/tube_geometry/tubularsegments' ) ).setWidth( '90px' ) );
	tubularSegmentsRow.add( tubularSegments );

	container.add( tubularSegmentsRow );

	// radialSegments

	const radialSegmentsRow = new UIRow();
	const radialSegments = new UIInteger( parameters.radialSegments ).onChange( update );

	radialSegmentsRow.add( new UIText( strings.getKey( 'sidebar/geometry/tube_geometry/radialsegments' ) ).setWidth( '90px' ) );
	radialSegmentsRow.add( radialSegments );

	container.add( radialSegmentsRow );

	// closed

	const closedRow = new UIRow();
	const closed = new UICheckbox( parameters.closed ).onChange( update );

	closedRow.add( new UIText( strings.getKey( 'sidebar/geometry/tube_geometry/closed' ) ).setWidth( '90px' ) );
	closedRow.add( closed );

	container.add( closedRow );

	// curveType

	const curveTypeRow = new UIRow();
	const curveType = new UISelect().setOptions( { centripetal: 'centripetal', chordal: 'chordal', catmullrom: 'catmullrom' } ).setValue( parameters.path.curveType ).onChange( update );

	curveTypeRow.add( new UIText( strings.getKey( 'sidebar/geometry/tube_geometry/curvetype' ) ).setWidth( '90px' ), curveType );

	container.add( curveTypeRow );

	// tension

	const tensionRow = new UIRow().setDisplay( curveType.getValue() == 'catmullrom' ? '' : 'none' );
	const tension = new UINumber( parameters.path.tension ).setStep( 0.01 ).onChange( update );

	tensionRow.add( new UIText( strings.getKey( 'sidebar/geometry/tube_geometry/tension' ) ).setWidth( '90px' ), tension );

	container.add( tensionRow );

	//

	function update() {

		tensionRow.setDisplay( curveType.getValue() == 'catmullrom' ? '' : 'none' );

		editor.execute( new SetGeometryCommand( editor, object, new THREE.TubeGeometry(
			new THREE.CatmullRomCurve3( points.getValue(), closed.getValue(), curveType.getValue(), tension.getValue() ),
			tubularSegments.getValue(),
			radius.getValue(),
			radialSegments.getValue(),
			closed.getValue()
		) ) );

	}

	return container;

}

export { GeometryParametersPanel };
