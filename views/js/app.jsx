import '../app.css'

class App extends React.Component {
  render() {
    return (
      <div className="main">
        <TextBox />
      </div>
    );
  }
}

class TextBox extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      initialLink: "",
      isSubmitted: false
    };

    this.textSubmit = this.textSubmit.bind(this);
  }

  textSubmit(event) {
    if (event.target.value != "") {
      this.setState({isSubmitted: true});
      this.setState({initialLink: event.target.value});
    } else {
      this.setState({isSubmitted: false});
    }
  }

  render() {
    let { isSubmitted } = this.state;
    let DisplayArea;

    if(isSubmitted) {
      DisplayArea = <h1>{this.state.initialLink}</h1>;
    } else {
      DisplayArea = <h1>Input enter the Url !</h1>;
    }

    return (
      <div className="textbox">
        {DisplayArea}
       
        <input type="text" onChange={this.textSubmit}/>
        <input type="submit" value="Submit"/>

      </div>
    );
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('app')
);