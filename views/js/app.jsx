import './app.css'
import TextBox from './textbox.jsx'

class App extends React.Component {
  render() {
    return (
      <TextBox placeholder={"Enter a Url"}/>
    );
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('app')
);