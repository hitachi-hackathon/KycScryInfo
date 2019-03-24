import React, {Component} from "react"
import { Container,Row,Col} from 'react-bootstrap'

class Home extends Component{
    render(){
        return(
            <div>
                <Container>
                    <Row>
                        <Col></Col>
                        <Col xs={12} md={12} lg={9} >
                            <h3 class="title1 text-center" >Welcome to faceX</h3>
                        </Col>
                        <Col></Col>
                    </Row>
                    <Row>
                        <Col lg={1} ></Col>
                        <Col xs={12} sm={12} lg={10}>
                            <p>Our solution</p>
                        </Col>
                        <Col lg={1} ></Col>
                    </Row>
                </Container>
            </div>
        )
    }
}

export default Home