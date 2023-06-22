import { LineSegments2 } from '/static/freecad/LineSegments2.js';
import { LineGeometry } from '/static/freecad/LineGeometry.js';
import { LineMaterial } from '/static/freecad/LineMaterial.js';

var Line2 = function ( geometry, material ) {

	if ( geometry === undefined ) geometry = new LineGeometry();
	if ( material === undefined ) material = new LineMaterial( { color: Math.random() * 0xffffff } );

	LineSegments2.call( this, geometry, material );

	this.type = 'Line2';

};

Line2.prototype = Object.assign( Object.create( LineSegments2.prototype ), {

	constructor: Line2,

	isLine2: true

} );

export { Line2 };
