import '../app.css'
import axios from 'axios'

class App extends React.Component {
  render() {
    return (
      <TextBox />
    );
  }
}

class TextBox extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      initialLink: "",
      shortenedLink: "",
      isSubmitted: false
    };

    this.textSubmit = this.textSubmit.bind(this);
    this.shortenUrl = this.shortenUrl.bind(this);
  }

  textSubmit(event) {
    if (event.target.value != "") {
      this.setState({initialLink: event.target.value});
    } 
  }

  shortenUrl() {
    axios.post('/create-short-url', {
      long_url: this.state.initialLink,
      user_id: "e0dba740-fc4b-4977-872c-d360239e6b1a"             
    })
    .then( (response) => {
      if (response.status == 200) {
        this.setState({shortenedLink: response.data['short_url']});
        this.setState({isSubmitted: true});
      } else {
        this.setState({isSubmitted: false});
      }
    })
    .catch((error) => console.log(error))
  }

  render() {
    let { isSubmitted } = this.state;
    let DisplayArea;

    if(isSubmitted) {
      DisplayArea = <h1 id='url'>{this.state.shortenedLink}</h1>;
    } else {
      DisplayArea = <h1 id='demo'>Input enter the Url !</h1>;
    }

    return (
      <div className="textbox">
        {DisplayArea}
       
        <input type="text" onChange={this.textSubmit} />
        <input type="submit" value="Submit" onClick={this.shortenUrl} />

      </div>
    );
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('app')
);