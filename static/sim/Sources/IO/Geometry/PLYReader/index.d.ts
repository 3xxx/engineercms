import { vtkAlgorithm, vtkObject } from "../../../interfaces";

interface IPLYReaderOptions {
	binary?: boolean;
	compression?: string;
	progressCallback?: any;
}

/**
 * 
 */
export interface IPLYReaderInitialValues {}

type vtkPLYReaderBase = vtkObject & Omit<vtkAlgorithm,
	| 'getInputData'
	| 'setInputData'
	| 'setInputConnection'
	| 'getInputConnection'
	| 'addInputConnection'
	| 'addInputData'>;

export interface vtkPLYReader extends vtkPLYReaderBase {

	/**
	 * 
	 */
	getBaseURL(): string;

	/**
	 * 
	 */
	getDataAccessHelper(): any;

	/**
	 * Get the url of the object to load.
	 */
	getUrl(): string;

	/**
	 * Load the object data.
	 * @param {IPLYReaderOptions} [options] 
	 */
	loadData(options?: IPLYReaderOptions): Promise<any>;

	/**
	 * Parse data.
	 * @param {String | ArrayBuffer} content The content to parse.
	 */
	parse(content: string | ArrayBuffer): void;

	/**
	 * Parse data as ArrayBuffer.
	 * @param {ArrayBuffer} content The content to parse. 
	 */
	parseAsArrayBuffer(content: ArrayBuffer): void;

	/**
	 * Parse data as text.
	 * @param {String} content The content to parse. 
	 */
	parseAsText(content: string): void;
	/**
	 *
	 * @param inData 
	 * @param outData 
	 */
	requestData(inData: any, outData: any): void;

	/**
	 * Set the url of the object to load.
	 * @param {String} url the url of the object to load.
	 * @param {IPLYReaderOptions} [option] The PLY reader options.
	 */
	setUrl(url: string, option?: IPLYReaderOptions): boolean;

	/**
	 * 
	 * @param dataAccessHelper 
	 */
	setDataAccessHelper(dataAccessHelper: any): boolean;
}

/**
 * Method used to decorate a given object (publicAPI+model) with vtkPLYReader characteristics.
 *
 * @param publicAPI object on which methods will be bounds (public)
 * @param model object on which data structure will be bounds (protected)
 * @param {IPLYReaderInitialValues} [initialValues] (default: {})
 */
export function extend(publicAPI: object, model: object, initialValues?: IPLYReaderInitialValues): void;

/**
 * Method used to create a new instance of vtkPLYReader
 * @param {IPLYReaderInitialValues} [initialValues] for pre-setting some of its content
 */
export function newInstance(initialValues?: IPLYReaderInitialValues): vtkPLYReader;


/**
 * vtkPLYReader is a source object that reads polygonal data in Stanford
 * University PLY file format (see http://graphics.stanford.edu/data/3Dscanrep).
 * It requires that the elements "vertex" and "face" are defined. The "vertex"
 * element must have the properties "x", "y", and "z". The "face" element must
 * have the property "vertex_indices" defined. Optionally, if the "face" element
 * has the properties "intensity" and/or the triplet "red", "green", "blue", and
 * optionally "alpha"; these are read and added as scalars to the output data.
 * If the "face" element has the property "texcoord" a new TCoords point array
 * is created and points are duplicated if they have 2 or more different texture
 * coordinates.
 */
export declare const vtkPLYReader: {
	newInstance: typeof newInstance;
	extend: typeof extend;
}
export default vtkPLYReader;
