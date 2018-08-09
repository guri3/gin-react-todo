import React from 'react'
import TodoList from './todo_list'
import Header from './header'
import AddTodo from './add_todo'

export default class Main extends React.Component{
  constructor(props){
    super(props)
    this.state = {
      data: []
    }
    this.addTodoComp = this.addTodoComp.bind(this)
  }

  componentWillMount() {
    fetch('http://localhost:3030/todos')
      .then(response => response.json())
      .then(data => {
        this.setState({data: data.result})
      })
  }

  addTodoComp(){
    fetch('http://localhost:3030/todos')
      .then(response => response.json())
      .then(data => {
        this.setState({data: data.result})
      })
  }

  render(){
    return(
      <div>
        <Header />
        <TodoList data={this.state.data} />
        <AddTodo addTodoHandler={this.addTodoComp}/>
      </div>
    )
  }
}
