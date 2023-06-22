import os,sys,argparse
from ansys.mapdl import reader as pymapdl_reader
import vtkmodules.all as vtk
# import vtk ！！！！坑1，必须用上面的，引入all，才能编译成exe后不缺少vtk包

# >>> geom.save('mesh.vtk')
#         Notes
#         -----
#         Binary files write much faster than ASCII and have a smaller
#         file size.
# Notes
#         -----
#         Binary files write much faster than ASCII, but binary files
#         written on one system may not be readable on other systems.
#         Binary can only be selected for the legacy writer.
#         Examples
#         --------
#         Write nodal results as a binary vtk file.
#         >>> rst.save_as_vtk('results.vtk')
#         Write using the xml writer
#         >>> rst.save_as_vtk('results.vtu')
#         Write only nodal and elastic strain for the first result
#         >>> rst.save_as_vtk('results.vtk', [0], ['EEL', 'EPL'])
#         Write only nodal results (i.e. displacements) for the first result.
#         >>> rst.save_as_vtk('results.vtk', [0], [])
#         """

def read_ansys(input_file,mid_file,output_file):
   try:
     result = pymapdl_reader.read_binary(input_file)   #新版读文件
     result.save_as_vtk(mid_file)   #这里写入binary格式

     reader = vtk.vtkUnstructuredGridReader()
     reader.SetFileName(mid_file)
     reader.Update()
     output = reader.GetOutput()
     geometry_filter = vtk.vtkGeometryFilter()
     geometry_filter.SetInputData(output)
     geometry_filter.Update()
     plyWriter = vtk.vtkXMLPolyDataWriter()
     plyWriter.SetFileName(output_file)
     plyWriter.SetInputConnection(geometry_filter.GetOutputPort())
     plyWriter.SetDataModeToBinary()    #写成binary格式。不要这句，则默认是xml格式
     plyWriter.Write()
     return "ok"
   except:
     # print('调用错误:')
     return "err"

if __name__ == '__main__':

    # parser.add_argument('-u', '--user', dest='User', type=str,default='root', help='target User')
    # parser.add_argument('-s', '--sex', dest='Sex', type=str, choices=['男', '女'], default='男', help='target Sex')
    # parser.add_argument('-n', '--number', dest='Num', nargs=2, required=True,type=int, help='target Two Numbers')

    parser = argparse.ArgumentParser(description="读取ansys计算结果文件(*.rst)转换成vtk(unstructured grid)图形文件和表面vtp(polydata)图形文件")
    parser.add_argument('-i', '--input_file', dest='input_file', required=True,type=str, help='target InputFile Ansys rst')
    parser.add_argument('-f', '--middle_file', dest='middle_file', required=True,type=str,  help='target MiddleFile vtk')
    parser.add_argument('-o', '--output_file', dest='output_file', required=True,type=str, help='target OutputFile vtp')

    # ！！！！坑2：-m 是绝对不行的，不能用这个字母，奇怪！！
    # dir_path = os.path.dirname(os.path.realpath(__file__))   #目录路径
    # input_file=os.path.join(dir_path, 'damrst.rst')
    # mid_file=os.path.join(dir_path, 'damrst.vtk')
    # output_file=os.path.join(dir_path, 'damrst.vtp')
    # sys.exit(read_ansys(input_file,mid_file,output_file))

    args = parser.parse_args()
    sys.exit(read_ansys(args.input_file,args.middle_file,args.output_file))
    # Python打包exe文件出现no module named ‘vtkmodules‘
    # https://blog.csdn.net/qq_41868425/article/details/112691364

    #  pyinstaller -F plot.py  （-F表示打包单个文件，-w是为了打开exe时候不弹出黑框）
    # ！！！！坑3！！！不要用-w 选项来编译，否则不出现命令行提示，也无法显示-h help内容。
    # 很显然没找到matplotlib.backends.backend_tkagg
    # 于是将打包命令后面加上如下：
    # D:\pythonWork\baiduai>pyinstaller -F machine1.py --hidden-import matplotlib.backends.backend_tkagg
