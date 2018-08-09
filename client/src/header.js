import React from 'react'
import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'

export default class Header extends React.Component{
  constructor(props){
    super(props)
    this.state = {
    }
  }

  render(){
    return(
      <div>
        <AppBar position="static">
          <Toolbar>
            <Typography variant="title" color="inherit">
              ToDo
            </Typography>
          </Toolbar>
        </AppBar>
      </div>
    )
  }
}
