import React, {Component} from "react"
import { Container,Row,Col} from 'react-bootstrap'
import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://192.168.43.5:8081',
  crossDomain: true
});

class Grant extends Component{
    constructor(props){
        super(props);
        this.state=({
            status:"0",
            address:"User address",
            dexAddress:"institute address",
            bank:"ICBC Bank"
        })

    }

    catchRes(res){
        // const msg = res.data;
        console.log(res);
        this.setState({
            status:"1",
        });
    }

    catchErr(error){
      console.log(error.response);
    }

    submit(){
        this.catchRes("fsdf");
        // instance.get('/user/grant',{
        //     address:this.state.address,
        //     institute_address:this.state.dexAddress,
        // }).then((res)=>{
        //     this.catchRes(res);
        // }).catch((err)=>{
        //     this.catchErr(err);
        // })
    }

    changeBank(e){
        this.setState({
            gender:e.target.value
        });
    }

    render(){
        let content;
        if(this.state.status === "0"){
            content = (
            <div class="row align-items-center h-100">
                <div class="col ">
                    <h3 class="title2 text-center" >Grant Institute your KYC certificate</h3>
                    <form id="form">
                        <div class="form-group">
                        <label for="exampleInputEmail1">Institutes</label>
                        <select class="form-control" value={this.state.bank} onChange={this.changeBank.bind(this)}>
                          <option value="ICBC Bank" >ICBC Bank</option>
                          <option value="Citi Bank" >Citi Bank</option>
                        </select>
                        </div>
                        <button type="submit" class="btn btn-primary" onClick={this.submit.bind(this)}
                        >Submit
                        </button>
                    </form>
                </div>
            </div>);
        }else{
            content = (
            <div class="row  align-items-center h-100">
                <div id="status" class="col text-center">
                    <div id="spinner" class="spinner-border text-primary" role="status">
                      <p><span class="sr-only">Loading...</span></p>
                    </div>
                    <h3 id="process" class="title2 text-center" >Your data has send to Institute for confirmation</h3>
                    <h3 class="title2 text-center" >Please wait</h3>
                </div>
            </div>)
        }
        return(
            <div>
                <Container>
                    <Row>
                        <Col></Col>
                        <Col xs={12} md={12} lg={9} >
                        {content}
                        </Col>
                        <Col></Col>
                    </Row>
                </Container>
            </div>
        )
    }
}

export default Grant