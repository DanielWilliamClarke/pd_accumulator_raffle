import React from 'react';
import ReactDOM from 'react-dom';
import Countdown from 'react-countdown-now';

require("./style.css");

class Root extends React.Component {

  constructor() {
    super();

    // Define the initial state:
    this.state = {
      round: "Awaiting",
      goldenRound: {
        next: false,
        now: false
      },
      selectedParticipant: null,
      scoreBoard: []
    };
  }

  update(updateObj) {
    this.setState(updateObj);
    // Force a render without state change...
    this.forceUpdate();
  }
  
  render() {
    
    let goldenRoundNextDom = ""
    if(this.state.goldenRound.next || this.state.goldenRound.now) {
      goldenRoundNextDom = <div className="goldenround">
        <h2 className="heading">
          Golden Round{!this.state.goldenRound.now ? " Imminent": ""}!!
        </h2>
      </div>
    }

    let selectedParticipantDom = <h1>Awaiting Winner</h1>
    if(this.state.selectedParticipant) {
      selectedParticipantDom = <div className={this.state.goldenRound.now ? "goldenround": "winner"}>
        {goldenRoundNextDom}
        <h2>
          Winner: {this.state.selectedParticipant.participant}! Score: {this.state.selectedParticipant.score}  
        </h2>
      </div>
    }

    let scoreBoardDom = <h1>Awaiting Scores</h1>
    if(this.state.scoreBoard.length) {
      scoreBoardDom = <table className="raffle_score_board">
        <tbody>
          <tr className="scoreheader">
            <th>Player</th>
            <th>Score</th>
          </tr>
          {this.state.scoreBoard.map((sb, i) => {
            return <tr className="scoreresults" key={i}>
              <td>{sb.participant}</td>
              <td>{sb.score}</td>
            </tr>
          })}
        </tbody>
      </table>    
    }

    return(
      <div>
        <h2 className="round">
          Round <u>{this.state.round}</u> <i>Next Draw In!</i>
        </h2>
        <h1 className="countdown">
          <div>
            <Countdown date={Date.now() + (1000 * 20)} /> 
          </div>
        </h1>
        {selectedParticipantDom}
        <hr />
        {scoreBoardDom}
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

    // event name
    console.log(obj.event);
    // event data
    console.log(obj.update);

    component.update(obj.update)
  }

  //Reload on keypress 'r'
document.addEventListener('keyup', function(e){
  if(e.keyCode == 82)
    window.location.reload();
  })
}


