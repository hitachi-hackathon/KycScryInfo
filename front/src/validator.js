import React, {Component} from "react"
import { Container,Row,Col} from 'react-bootstrap'
import axios from "axios"

const instance = axios.create({
  baseURL: 'http://172.16.0.38:8081',
  crossDomain: true
});

class Certify extends Component{
    constructor(props){
        super(props);
        this.state=({
            address:"0xdffdf",
            user_address:"",
            name:"",
            gender:"",
            country:"",
            age:"",
            residency_address:"",
            status:""
        })
    }

    catchRes(res){
        const msg = res.data;
        console.log(res);
        this.setState({
            user_address:msg.address,
            name:msg.name,
            gender:msg.dender,
            country:msg.country,
            age:msg.country,
            residency_address:msg.residency_address,
            status:"pending"
        });
    }

    catchRes2(res){
        const msg = res.data;
        console.log(res);
        const status = msg.status? "Passed" :"Failed"
        this.setState({
            status:status
        });
    }

    catchErr(error){
      console.log(error.response);
    }

    componentDidMount(){
        instance.get('/authority/users',{
            address:this.state.address
        }) // get asic data
        .then((res)=>{
        this.catchRes(res);
        }).catch((err)=>{
        this.catchErr(err);
        });
    }

    submit(){
        instance.post('/authority/certify',{
            address:this.state.address,
            user_address:this.state.user_address
        }) // get asic data
        .then((res)=>{
        this.catchRes2(res);
        }).catch((err)=>{
        this.catchErr(err);
        });
    }

    render(){
        return(
            <div>
                <Container>
                    <Row>
                        <Col></Col>
                        <Col xs={12} md={12} lg={9} >
                            <h3 class="title1 text-center" >Welcome to faceX, Smart KYC</h3>
                        </Col>
                        <Col></Col>
                    </Row>
                    <Row>
                        <Col lg={1} ></Col>
                        <Col xs={12} sm={12} lg={10}>
                            <table class="table">
                              <thead>
                                <tr>
                                    <th scope="col">User</th>
                                    <th scope="col">Gender</th>
                                    <th scope="col">Country</th>
                                    <th scope="col">Age</th>
                                    <th scope="col">Address</th>
                                    <th scope="col">Account Address</th>
                                    <th scope="col">Grant Access</th>
                                    <th scope="col">Status</th>
                                </tr>
                              </thead>
                              <tbody>
                                <tr>
                                  <th scope="row">{this.state.name}</th>
                                  <td>{this.state.gender}</td>
                                  <td>{this.state.country}</td>
                                  <td>{this.state.age}</td>
                                  <td>{this.state.residency_address}</td>
                                  <td>{this.state.user_address}</td>
                                  <td>
                                    {this.state.address? <button type="submit" class="btn btn-primary"
                                        onClick={this.submit.bind(this)}
                                    >Grant Certificate</button>
                                    :<div></div>
                                    }
                                  </td>
                                  <td>{this.state.status}</td>
                                </tr>
                              </tbody>
                            </table>
                        </Col>
                        <Col lg={1} ></Col>
                    </Row>
                </Container>
            </div>
        )
    }
}

export default Certify