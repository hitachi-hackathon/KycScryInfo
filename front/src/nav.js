// packages
import React, { Component } from 'react'
import { Link, NavLink } from 'react-router-dom'
import { Navbar,Nav,NavDropdown} from 'react-bootstrap'
// components


class NavBar extends Component{
    render(){
        return(
            <div>
                <Navbar id="navbar" bg="light" expand="md">
                    <Navbar.Brand as={Link} to="/" >FaceX</Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav"/>
                    <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="mr-auto">
                        <NavDropdown title="User" id="basic-nav-dropdown">
                            <NavDropdown.Item as={Link} to="/user/verify">Verify</NavDropdown.Item>
                            <NavDropdown.Item as={Link} to="/user/grant">Grant Access</NavDropdown.Item>
                            <NavDropdown.Divider />
                            <NavDropdown.Item as={Link} to="/user/list">KYC List</NavDropdown.Item>
                        </NavDropdown>
                        <Nav.Link as={Link} to="/institute" >Institute</Nav.Link>
                        <NavDropdown title="Authority" id="basic-nav-dropdown">
                            <NavDropdown.Item as={Link} to="/authority/certify">Certify</NavDropdown.Item>
                            <NavDropdown.Item as={Link} to="/authority/verify">Verify</NavDropdown.Item>
                        </NavDropdown>
                    </Nav>
                    <Nav.Link as={Link} to="/about" >About</Nav.Link>
                </Navbar.Collapse>
                </Navbar>
            </div>
    )}
}

export default NavBar;