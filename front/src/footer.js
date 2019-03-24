// packages
import React, { Component } from 'react';
import {Link, withRouter} from 'react-router-dom';
import { Container,Row,Col} from 'react-bootstrap';
// components
import Logo from "./pic/finger.png"
import email from './pic/email.svg';
import telegram from './pic/telegram.svg';

class Footer extends Component{
    render(){
        return(
            <footer id="footer">
                <div id="footer-up">
                    <Container>
                        <Row>
                            <Col></Col>
                            <Col xs={12} md={12} lg={9}>
                                <Row>
                                <Col md={3} sm={4} xs={12} ><p class="footer-title">FaceX</p></Col>
                                <Col md={3} sm={3} xs={12} ></Col>
                                <Col md={3} sm={3} xs={6} >
                                    <p class="footer-title">Solution</p>
                                    <Link class="footer-link" to="/user" >User</Link><br/>
                                    <Link class="footer-link" to="/institute" >Institute</Link><br/>
                                    <Link class="footer-link" to="/validator" >Validator</Link>
                                </Col>
                                <Col md={3} sm={2} xs={6} >
                                    <p class="footer-title">About</p>
                                    <Link class="footer-link" to="/about#company" >Company</Link><br/>
                                </Col>
                                </Row>
                            </Col>
                            <Col></Col>
                        </Row>
                    </Container>
                </div>
                <div id="footer-down">
                    <Container >
                        <Row className="align-items-center">
                            <Col>© 2019 faceX</Col>
                            <Col>
                                <ul class="nav d-flex flex-row-reverse"  id="footer-icons">
                                <li class="nav-item">
                                <a class="nav-link" href="mailto:cryptoattackio@gmail.com">
                                <img class="footer-icon "src={email} />
                                </a></li>
                                <li class="nav-item">
                                <a class="nav-link" href="https://t.me/joinchat/G0CI1BKgFELN9FLcOfrsLQ">
                                <img class="footer-icon "src={telegram} />
                                </a></li>
                                
                                </ul>
                            </Col>
                        </Row>
                    </Container>
                </div>
            </footer>
        )
    }
}

export default withRouter(Footer);