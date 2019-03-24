import React, {Component} from "react"
import { Container,Row,Col} from 'react-bootstrap'
import axios from 'axios';
const instance = axios.create({
  baseURL: 'http://172.16.0.38:8081',
  crossDomain: true
});

class Institute extends Component{
    constructor(props){
        super(props);
        this.state=({
            dexAddress:"0xdffdf",
            user:"",
            status:""
        })
    }

    catchRes(res){
        const msg = res.data;
        console.log(res);
        this.setState({
            user:msg.user,
            status:msg.status
        });
    }

    catchErr(error){
      console.log(error.response);
    }

    componentDidMount(){
        // send request to server and get data.
        // console.log("request data");
        instance.get('/institute',{
            address:this.state.dexAddress
        }) // get asic data
        .then((res)=>{
        this.catchRes(res);
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
                            <h3 class="title1 text-center" >The status of KYC verification</h3>
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
                                  <th scope="col">Status</th>
                                </tr>
                              </thead>
                              <tbody>
                                <tr>
                                  <th scope="row">{this.state.user}</th>
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

export default Institute