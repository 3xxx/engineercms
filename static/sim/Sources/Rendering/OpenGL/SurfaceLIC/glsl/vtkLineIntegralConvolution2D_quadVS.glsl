//VTK::System::Dec

/*=========================================================================

  Program:   Visualization Toolkit
  Module:    vtktextureObjectVS.glsl

  Copyright (c) Ken Martin, Will Schroeder, Bill Lorensen
  All rights reserved.
  See Copyright.txt or http://www.kitware.com/Copyright.htm for details.

     This software is distributed WITHOUT ANY WARRANTY; without even
     the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR
     PURPOSE.  See the above copyright notice for more information.

=========================================================================*/

attribute vec4 vertexDC;
attribute vec2 tcoordDC;
varying vec2 tcoordVC;

void main()
{
  tcoordVC = tcoordDC;
  gl_Position = vertexDC;
}
