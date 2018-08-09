import React from 'react'
import Todo from './todo';

export default class TodoList extends React.Component{
  constructor(props){
    super(props)
    this.state = {
    }
  }

  render(){
    let todoList = []
    for(let i in this.props.data){
      todoList.push(<div><Todo data={this.props.data[i]}/></div>)
    }

    return(
      <div>
        {todoList}
      </div>
    )
  }
}
