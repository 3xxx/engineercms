export default [
  {
    label: 'Choose a headset',
  },
  {
    id: 'CardboardV1',
    label: 'Cardboard I/O 2014',
    fov: 40,
    interLensDistance: 0.06,
    baselineLensDistance: 0.035,
    screenLensDistance: 0.042,
    distortionCoefficients: [0.441, 0.156],
    inverseCoefficients: [
      -0.4410035, 0.42756155, -0.4804439, 0.5460139, -0.58821183, 0.5733938,
      -0.48303202, 0.33299083, -0.17573841, 0.0651772, -0.01488963, 0.001559834,
    ],
  },
  {
    id: 'CardboardV2',
    label: 'Cardboard I/O 2015',
    fov: 60,
    interLensDistance: 0.064,
    baselineLensDistance: 0.035,
    screenLensDistance: 0.039,
    distortionCoefficients: [0.34, 0.55],
    inverseCoefficients: [
      -0.33836704, -0.18162185, 0.862655, -1.2462051, 1.0560602, -0.58208317,
      0.21609078, -0.05444823, 0.009177956, -9.904169e-4, 6.183535e-5,
      -1.6981803e-6,
    ],
  },
];
