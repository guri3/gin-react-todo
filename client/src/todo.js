import React from 'react'
import Checkbox from '@material-ui/core/Checkbox'

export default class Todo extends React.Component{
  constructor(props){
    super(props)
    this.state = {
      done: this.props.data.done
    }
    this.clickCheckBox = this.clickCheckBox.bind(this)
  }

  clickCheckBox(){
    this.setState({done: !this.state.done})
    const url = 'http://localhost:3030/todos/' + this.props.data.id + "/toggle"
    fetch(url, {
      method: 'POST'
    })
  }

  render(){
    const done = this.props.data.done ? '完了' : '未完了'
    return(
      <div>
        <Checkbox
          checked={this.state.done}
          onClick={this.clickCheckBox}
        />
        {this.props.data.name}
      </div>
    )
  }
}
