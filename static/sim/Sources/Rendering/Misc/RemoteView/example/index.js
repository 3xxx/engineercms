import vtkWSLinkClient from 'vtk.js/Sources/IO/Core/WSLinkClient';
import SmartConnect from 'wslink/src/SmartConnect';
import vtkRemoteView, {
  connectImageStream,
} from 'vtk.js/Sources/Rendering/Misc/RemoteView';

vtkWSLinkClient.setSmartConnectClass(SmartConnect);

document.body.style.padding = '0';
document.body.style.margin = '0';

const divRenderer = document.createElement('div');
document.body.appendChild(divRenderer);

divRenderer.style.position = 'relative';
divRenderer.style.width = '100vw';
divRenderer.style.height = '100vh';
divRenderer.style.overflow = 'hidden';

const view = vtkRemoteView.newInstance({
  rpcWheelEvent: 'viewport.mouse.zoom.wheel',
});
view.setContainer(divRenderer);
view.setInteractiveRatio(0.7); // the scaled image compared to the clients view resolution
view.setInteractiveQuality(50); // jpeg quality

window.addEventListener('resize', view.resize);

const clientToConnect = vtkWSLinkClient.newInstance();

// Error
clientToConnect.onConnectionError((httpReq) => {
  const message =
    (httpReq && httpReq.response && httpReq.response.error) ||
    `Connection error`;
  console.error(message);
  console.log(httpReq);
});

// Close
clientToConnect.onConnectionClose((httpReq) => {
  const message =
    (httpReq && httpReq.response && httpReq.response.error) ||
    `Connection close`;
  console.error(message);
  console.log(httpReq);
});

// hint: if you use the launcher.py and ws-proxy just leave out sessionURL
// (it will be provided by the launcher)
const config = {
  application: 'cone',
  sessionURL: 'ws://localhost:1234/ws',
};

// Connect
clientToConnect
  .connect(config)
  .then((validClient) => {
    connectImageStream(validClient.getConnection().getSession());

    const session = validClient.getConnection().getSession();
    view.setSession(session);
    view.setViewId(-1);
    view.render();
  })
  .catch((error) => {
    console.error(error);
  });
