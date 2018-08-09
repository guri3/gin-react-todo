import React from 'react'
import TextField from '@material-ui/core/TextField'
import Button from '@material-ui/core/Button'
import AddIcon from '@material-ui/icons/Add'

export default class AddTodo extends React.Component{
  constructor(props) {
    super(props)
    this.state = {
      name: ''
    }
    this.onChangeHandler = this.onChangeHandler.bind(this)
    this.buttonClickHandler = this.buttonClickHandler.bind(this)
    this.sendCommand = this.sendCommand.bind(this)
  }

  onChangeHandler(e){
    this.setState({name: e.target.value})
  }

  buttonClickHandler(){
    // ボタンを押した時の処理を記述
    const obj = {name: this.state.name}
    const method = 'POST'
    const body = JSON.stringify(obj)
    fetch('http://localhost:3030/todos', {
      method: method,
      body: body,
    })
    .then((response) => {
      response.json()
    })
    .then((data) => console.log(data))
    this.setState({name: ''})
    this.props.addTodoHandler()
  }

  sendCommand(e){
    const ENTER = 13;
    if (e.keyCode !== ENTER) {
      return
    }
    // ボタンを押した時の処理を記述
    const obj = {name: this.state.name}
    const method = 'POST'
    const body = JSON.stringify(obj)
    fetch('http://localhost:3030/todos', {
      method: method,
      body: body,
    })
    .then((response) => {
      response.json()
    })
    .then((data) => console.log(data))
    this.setState({name: ''})
    this.props.addTodoHandler()
  }

  render() {
    return (
      <div>
        <TextField
          onChange={this.onChangeHandler}
          onKeyDown={this.sendCommand}
          value={this.state.name}
        />
        <Button
          variant="fab"
          mini
          color="primary"
          aria-label="Add"
          onClick={this.buttonClickHandler}>
          <AddIcon />
        </Button>
      </div>
    )
  }
}
