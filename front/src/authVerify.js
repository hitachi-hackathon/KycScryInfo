import React, {Component} from "react"
import { Container,Row,Col} from 'react-bootstrap'
import axios from "axios"

const instance = axios.create({
  baseURL: 'http://172.16.0.38:8081',
  crossDomain: true
});

class Verify extends Component{
    constructor(props){
        super(props);
        this.state=({
            address:"0xdffdf",
            institute_address:"",
            user_address:"",
            status:""
        })
    }

    catchErr(error){
      console.log(error.response);
    }

    catchRes(res){
        const msg = res.data;
        console.log(res);
        this.setState({
            user_address:msg.user_address,
            institute_address:msg.institute_address
        });
    }

    componentDidMount(){
        instance.get('/authority/institutes',{
            address:this.state.address
        }) // get asic data
        .then((res)=>{
        this.catchRes(res);
        }).catch((err)=>{
        this.catchErr(err);
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

    submit(){
        instance.post('/authority/verify',{
            authority_address:this.state.address,
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
                            <h3 class="title2 text-center" >KYC to Verify</h3>
                            <table class="table">
                              <thead>
                                <tr>
                                    <th scope="col">Institute</th>
                                    <th scope="col">User</th>
                                    <th scope="col">Verify</th>
                                    <th scope="col">Status</th>
                                </tr>
                              </thead>
                              <tbody>
                                <tr>
                                  <th scope="row">{this.state.institute_address}</th>
                                  <td>{this.state.user_address}</td>
                                  <td>
                                    {this.state.address?<button type="submit" class="btn btn-primary"
                                        onClick={this.submit.bind(this)}
                                    >Verify User</button>
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

export default Verify