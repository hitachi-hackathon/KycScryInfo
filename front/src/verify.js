import React, {Component} from "react"
import { Container,Row,Col} from 'react-bootstrap'
import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://172.16.0.38:8081',
  crossDomain: true
});

class Verify extends Component{
    constructor(props){
        super(props);
        this.state=({
            address:"0x3c4d26e916d79fc3fc925027a79612012462f691",
            name:"",
            gender:"Male",
            country:"",
            age:"",
            residency_address:"",
            document:null,
            state:"0"
        })

    }
    changeName(event){
        this.setState({
            name:event.target.value
        });
    }

    changeGender(e){
        this.setState({
            gender:e.target.value
        });
    }

    changeCountry(e){
        this.setState({
            country:e.target.value
        });
    }

    changeAge(e){
        this.setState({
            age:e.target.value
        });
    }

    changeAddress(e){
        this.setState({
            residency_address:e.target.value
        });
    }

    changeFile(e){
        this.setState({
            document:e.target.value
        });
    }

    catchRes(res){
        const msg = res.data;
        console.log(res);
        this.setState({
            state:"2",
        });
    }

    catchErr(error){
      console.log(error.response);
    }

    onWait(){
        console.log("waiting")
        this.setState({
            state:"1"
        })
    }
    wait(ms){
       var start = new Date().getTime();
       var end = start;
       while(end < start + ms) {
         end = new Date().getTime();
      }
    }

    send(){
        console.log("execute");
        instance.post('/user/upload',{
            address:this.state.address,
            name:this.state.name,
            gender:this.state.gender,
            country:this.state.country,
            age:this.state.age,
            residency_address:this.state.residency_address,
        }).then((res)=>{
            this.catchRes(res);
        }).catch((err)=>{
            this.catchErr(err);
        })
    }


    submit(){
        this.setState({
            state:"1"
        },function(){
            // this.forceUpdate();
            // this.wait(7000);
            // this.catchRes("fsd");
            this.send(); 
        })
         
    }



    render(){
        let content;
        if(this.state.state === "0"){
            content = (
                <div>
                    <h3 class="title2 text-center" >Please provide us your information</h3>
                    <form id="form">
                        <div class="form-group">
                        <label for="exampleInputEmail1">Name</label>
                        <input class="form-control" placeholder="Your name" 
                            onChange={this.changeName.bind(this)}
                        />
                        </div>
                        <div class="form-group">
                        <label for="exampleInputEmail1">Gender</label>
                        <select class="form-control" value={this.state.gender} onChange={this.changeGender.bind(this)}>
                          <option value="Male" >Male</option>
                          <option value="Female" >Female</option>
                        </select>
                        </div>
                        <div class="form-group">
                        <label for="exampleInputEmail1">Country</label>
                        <input class="form-control" placeholder="Country of Residency" 
                            onChange={this.changeCountry.bind(this)}
                        />
                        </div>
                        <div class="form-group">
                        <label for="exampleInputEmail1">Age</label>
                        <input class="form-control" placeholder="your age" 
                            onChange={this.changeAge.bind(this)}
                        />
                        </div>
                        <div class="form-group">
                        <label for="exampleInputEmail1">Address</label>
                        <input class="form-control" placeholder="" 
                            onChange={this.changeAddress.bind(this)}
                        />
                        </div>
                        <div class="form-group">
                        <label for="exampleFormControlFile1">Upload your Passport</label>
                        <input type="file" class="form-control-file"
                            onChange={this.changeFile.bind(this)}
                        />
                        </div>
                        <button type="submit" class="btn btn-primary"
                            onClick={this.submit.bind(this)}
                        >Submit</button>
                    </form>
                </div>
        )}else if(this.state.state === "1"){
            content = (
            <div class="row  align-items-center h-100">
                <div id="status" class="col text-center">
                    <div id="spinner" class="spinner-border text-primary" role="status">
                      <p><span class="sr-only">Loading...</span></p>
                    </div>
                    <h3 id="process" class="title2 text-center" >Your data is being processed</h3>
                    <h3 class="title2 text-center" >Please wait</h3>
                </div>
            </div>)
        }else{
            content = (
                <div class="row  align-items-center h-100">
                    <div class="col ">
                        <h3 class="title2 text-center" >Your File has successfully uploaded!</h3>
                        <h3 class="title2 text-center" >Please wait for KYC and AML check</h3>
                    </div>
                </div>
            )
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

export default Verify