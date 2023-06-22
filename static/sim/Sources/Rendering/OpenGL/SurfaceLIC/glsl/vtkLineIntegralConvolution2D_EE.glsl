//VTK::System::Dec

//=========================================================================
//
//  Program:   Visualization Toolkit
//  Module:    vtkLineIntegralConvolution2D_fs2.glsl
//
//  Copyright (c) Ken Martin, Will Schroeder, Bill Lorensen
//  All rights reserved.
//  See Copyright.txt or http://www.kitware.com/Copyright.htm for details.
//
//     This software is distributed WITHOUT ANY WARRANTY; without even
//     the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR
//     PURPOSE.  See the above copyright notice for more information.
//
//=========================================================================

// high-pass filter stage employed by vtkLineIntegralConvolution2D
// between LIC pass 1 and LIC pass 2. filtered LIC pass 1, becomes
// noise for pass2.

// the output of this shader
layout(location = 0) out vec4 EEOutput;

uniform sampler2D texLIC; // most recent lic pass
uniform float     uDx;    // fragment size
uniform float     uDy;    // fragment size

in vec2 tcoordVC;

// kernel for simple laplace edge enhancement.
// p=Laplace(p)+p
float K[9] = float[9](
  -1.0, -1.0, -1.0,
  -1.0,  9.0, -1.0,
  -1.0, -1.0, -1.0
  );

// determine if the fragment was masked
bool Masked(float val) { return val != 0.0; }

void main(void)
{
  // tex coord neighbor offsets
  vec2 fragDx[9] = vec2[9](
    vec2(-uDx, uDy), vec2(0.0, uDy), vec2(uDx, uDy),
    vec2(-uDx, 0.0), vec2(0.0, 0.0), vec2(uDx, 0.0),
    vec2(-uDx,-uDy), vec2(0.0,-uDy), vec2(uDx,-uDy)
    );

  vec2 lictc = tcoordVC.st;

  // compute the convolution but don't use convovled values if
  // any masked fragments on the stencil. Fragments outside
  // the valid domain are masked during initialization, and
  // texture wrap parameters are clamp to border with border
  // color that contains masked flag
  float conv = 0.0;
  bool dontUse = false;
  for (int i=0; i<9; ++i)
    {
    vec2 tc = lictc + fragDx[i];
    vec4 lic = texture2D(texLIC, tc);
    dontUse = dontUse || Masked(lic.g);
    conv = conv + K[i] * lic.r;
    }

  if (dontUse)
    {
    EEOutput = vec4(texture2D(texLIC, lictc).rg, 0.0, 1.0);
    }
  else
    {
    conv = clamp(conv, 0.0, 1.0);
    EEOutput = vec4(conv,texture2D(texLIC, lictc).g, 0.0, 1.0);
    }

}
