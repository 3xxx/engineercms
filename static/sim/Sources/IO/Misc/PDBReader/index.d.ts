import { vtkAlgorithm, vtkObject, vtkSubscription } from "../../../interfaces";


interface IPDBReaderOptions {
	binary?: boolean;
	compression?: string;
	progressCallback?: any;
}

/**
 * 
 */
export interface IPDBReaderInitialValues {
	numberOfAtoms?: number;
	requestCount?: number;
}

type vtkPDBReaderBase = vtkObject & Omit<vtkAlgorithm,
	| 'getInputData'
	| 'setInputData'
	| 'setInputConnection'
	| 'getInputConnection'
	| 'addInputConnection'
	| 'addInputData'>;

export interface vtkPDBReader extends vtkPDBReaderBase {

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
	 * 
	 */
	getNumberOfAtoms(): number;

	/**
	 * 
	 */
	getRequestCount(): number;

	/**
	 * 
	 * @param {Boolean} busy 
	 */
	invokeBusy(busy: boolean): void;

	/**
	 * 
	 */
	isBusy(): number;

	/**
	 * Load the object data.
	 * @param {IPDBReaderOptions} [options] 
	 */
	loadData(options?: IPDBReaderOptions): Promise<any>;

	/**
	 * 
	 * @param busy 
	 */
	onBusy(busy: boolean): vtkSubscription;

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
	 * 
	 * @param dataAccessHelper 
	 */
	setDataAccessHelper(dataAccessHelper: any): boolean;

	/**
	 * Set the url of the object to load.
	 * @param {String} url the url of the object to load.
	 * @param {IPDBReaderOptions} [option] The PDB reader options.
	 */
	setUrl(url: string, option?: IPDBReaderOptions): boolean;
}

/**
 * Method used to decorate a given object (publicAPI+model) with vtkPDBReader characteristics.
 *
 * @param publicAPI object on which methods will be bounds (public)
 * @param model object on which data structure will be bounds (protected)
 * @param {IPDBReaderInitialValues} [initialValues] (default: {})
 */
export function extend(publicAPI: object, model: object, initialValues?: IPDBReaderInitialValues): void;

/**
 * Method used to create a new instance of vtkPDBReader
 * @param {IPDBReaderInitialValues} [initialValues] for pre-setting some of its content
 */
export function newInstance(initialValues?: IPDBReaderInitialValues): vtkPDBReader;


/**
 * vtkPDBReader is a source object that reads Molecule files.
 */
export declare const vtkPDBReader: {
	newInstance: typeof newInstance;
	extend: typeof extend;
}
export default vtkPDBReader;
