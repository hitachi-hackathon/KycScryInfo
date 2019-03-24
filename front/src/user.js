import React, {Component} from "react"
import { Container,Row,Col} from 'react-bootstrap'
import {
  Route,
  Switch
} from 'react-router-dom';

import Verify from "./verify"
import Grant from "./grant"
import List from "./list"

class User extends Component{
    render(){
        return(
            <div>
                <Container>
                    <Row>
                        <Col></Col>
                        <Col xs={12} md={12} lg={9} >
                            <h3 class="title1 text-center" >Welcome, User Bob!</h3>
                        </Col>
                        <Col></Col>
                    </Row>
                </Container>

                <Route path="/user/verify" component={Verify} />
                <Route path="/user/grant" component={Grant} />
                <Route path="/user/list" component={List} />
            </div>
        )
    }
}

export default User