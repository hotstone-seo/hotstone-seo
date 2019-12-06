import React, { Component } from 'react';
import {

  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Col,
  Form,
  FormGroup,
   
  Label,
  Row,
} from 'reactstrap';

class RuleDetail extends Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.toggleFade = this.toggleFade.bind(this);
    this.state = {
      collapse: true,
      fadeIn: true,
      timeout: 300
    };
    this.handleBack = this.handleBack.bind(this);
  }

  toggle() {
    this.setState({ collapse: !this.state.collapse });
  }

  toggleFade() {
    this.setState((prevState) => { return { fadeIn: !prevState } });
  }
  handleBack() {

  }
  render() {
    const rule = this.props.match.params.id;

    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" md="9" lg="6">
            <Card>
              <CardHeader>
                <strong>Detail Rule ID {this.props.match.params.id}</strong>
              </CardHeader>
              <CardBody>
                <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Name</Label>
                    </Col>
                    <Col xs="12" md="9">
                       
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">URL Pattern</Label>
                    </Col>
                    <Col xs="12" md="9">
                      
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="textarea-input">Exclusion</Label>
                    </Col>
                    <Col xs="12" md="9">
                     
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Data Source</Label>
                    </Col>
                    <Col xs="12" md="9">
                      
                    </Col>
                  </FormGroup>
                </Form>
              </CardBody>
              <CardFooter>
              
                <Button type="button" size="md" color="secondary" onClick={this.handleBack}> Back</Button>
              </CardFooter>
            </Card>
          </Col>
        </Row>
      </div>
    );
  }
}

export default RuleDetail;
