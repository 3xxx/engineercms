import vtk
import os,sys
def unstructured_grid2poly(input_file, output_file):
    #reader = vtk.vtkXMLUnstructuredGridReader()
    reader = vtk.vtkUnstructuredGridReader()
    reader.SetFileName(input_file)
    print(reader)
    reader.Update()
    print(reader)
    output = reader.GetOutput()
   # print(output)
    geometry_filter = vtk.vtkGeometryFilter()
    geometry_filter.SetInputData(output)
    
    geometry_filter.Update()
    
    poly_out = geometry_filter.GetOutput()
    
    n_points = poly_out.GetNumberOfPoints()
    
    print('total number of poins {:g}'.format(n_points))
    
    plyWriter = vtk.vtkXMLPolyDataWriter()
    #plyWriter = vtk.vtkPolyDataWriter()
    plyWriter.SetFileName(output_file)
    plyWriter.SetInputConnection(geometry_filter.GetOutputPort())
    #plyWriter.SetInputConnection(geometry_filter.GetOutput())
    plyWriter.Write()

if __name__ == '__main__':

 dir_path = os.path.dirname(os.path.realpath(__file__))   #目录路径
 #input_file=os.path.join(dir_path, 'hex.vtk') 
 input_file=os.path.join(dir_path, 'damrst.vtk') 
 output_file=os.path.join(dir_path, 'damrst.vtp') 
 sys.exit(unstructured_grid2poly(input_file,output_file ))
