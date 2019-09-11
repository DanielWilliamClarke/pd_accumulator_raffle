var React = require('react');
var ReactDOM = require('react-dom');

require("./style.css");

class Root extends React.Component {

  constructor() {
    super();

    // Define the initial state:
    this.state = {
      backendText: false
    };
  }

  updateBackendText(message) {
    this.setState({
      backendText: message
    });
  }
  
  render() {
    return(
      <div>
        <h1 className="topic">
          Hello Gotron / React
        </h1>
        <h2 className="topic">
          {this.state.backendText}
        </h2>
      </div>
    )
  }
}

const ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");

window.onload = function () {
  var div = document.createElement("div");
  var component =  ReactDOM.render(<Root />, div);
  document.body.appendChild(div);

  ws.onmessage = (message) => {
    let obj = JSON.parse(message.data);

    component.updateBackendText(obj.AtrNameInFrontend)
    // event name
    console.log(obj.event);
    // event data
    console.log(obj.AtrNameInFrontend);
  }

  //Reload on keypress 'r'
  document.addEventListener('keyup', function(e){
    if(e.keyCode == 82)
      ws.send(JSON.stringify({
        "event": "hello-back",
        "AtrNameInFrontend": "Hello backend!",
      }))
    })
}


