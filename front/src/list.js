import React, {Component} from "react"
import { Container,Row,Col} from 'react-bootstrap'
import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://192.168.43.5:8081',
  crossDomain: true
});

class List extends Component{
    constructor(props){
        super(props);
        this.state=({
            address:"0xdffdf",
            institute:"",
            status:""
        })
    }

    catchRes(res){
        const msg = res.data;
        console.log(res);
        this.setState({
            institute:"ICBC Bank",
            status:msg.status
        });
    }

    catchErr(error){
      console.log(error.response);
    }

    componentDidMount(){
        // send request to server and get data.
        // console.log("request data");
        instance.get('/user/status',{
            address:this.state.address
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
                                  <th scope="col">Institutes</th>
                                  <th scope="col">Status</th>
                                </tr>
                              </thead>
                              <tbody>
                                <tr>
                                  <th scope="row">{this.state.institute}</th>
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

export default List