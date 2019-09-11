import React from 'react';
import ReactDOM from 'react-dom';
import Countdown from 'react-countdown-now';

require("./style.css");

class Root extends React.Component {

  constructor() {
    super();

    // Define the initial state:
    this.state = {
      round: 0,
      selectedParticipant: null,
      scoreBoard: []
    };
  }

  updateBackendText(round, selectedParticipant, scoreBoard) {
    this.setState({
      round: round,
      selectedParticipant: selectedParticipant,
      scoreBoard: scoreBoard
    });
  }
  
  render() {

    let scoreBoardDom = <h1>Awaiting Scores</h1>
    if(this.state.scoreBoard.length) {
      scoreBoardDom = this.state.scoreBoard.map(sb => {
        return <li>
          <span>
            {sb.participant}
          </span>
          <span>
            {sb.score}
          </span>
          </li>
      })
    }
    
    let selectedParticipantDom = <h1>Awaiting Winner</h1>
    if(this.state.selectedParticipant) {
      selectedParticipantDom = <div className="winner">
        <h2 className="heading">
          Round {this.state.round}
        </h2>
        <h2 className="heading">
          Winner: {this.state.selectedParticipant.participant}! 
        </h2>
        <h3 className="heading">
          Score: {this.state.selectedParticipant.score}
        </h3>
      </div>
    }

    return(
      <div>
        <h1 className="countdown">
          <div>
            <Countdown date={Date.now() + (1000 * 20)} /> 
          </div>
          <div>
            Until Next Draw!
          </div>
        </h1>

        {selectedParticipantDom}
        
        <h1>Raffle Scores:</h1>
        <ul className="raffle_score_board">
          <li><span>Participant</span><span>Score</span></li>
          {scoreBoardDom}
        </ul>
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

    component.updateBackendText(
      obj.update.round,
      obj.update.selectedParticipant,
      obj.update.scoreBoard)
  }
}


