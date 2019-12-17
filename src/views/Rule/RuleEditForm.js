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
  Input,
  Label,
  Row,
} from 'reactstrap';

class RuleEditForm extends Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.toggleFade = this.toggleFade.bind(this);
    this.state = {
      collapse: true,
      fadeIn: true,
      timeout: 300
    };
    this.handleCancel = this.handleCancel.bind(this);
  }

  toggle() {
    this.setState({ collapse: !this.state.collapse });
  }

  toggleFade() {
    this.setState((prevState) => { return { fadeIn: !prevState } });
  }

  handleCancel() {
    const { history } = this.props;
    history.push('/rule');
  }

  render() {
    const { data } = this.props.location
     
    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" md="9" lg="6">
            <Card>
              <CardHeader>
                <strong>Edit Rule</strong>
              </CardHeader>
              <CardBody>
                <Form method="post" className="form-horizontal">
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Name</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="hidden" id="id" name="id" value={data.id} />
                      <Input type="text" id="name" name="name" placeholder="Name" value={data.name}/>
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">URL Pattern</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="text" id="urlPattern" name="urlPattern" placeholder="URL Pattern" value={data.url_pattern} />
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Data Source</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="select" name="datasource" id="datasource">
                        <option value="1">Airport</option>
                      </Input>                      
                    </Col>
                  </FormGroup>
                </Form>
              </CardBody>
              <CardFooter>
              <Button type="submit" size="md" color="primary" style={{ marginRight: "0.4em" }}><i className="fa fa-dot-circle-o"></i> Save Change</Button>
                <Button type="button" size="md" color="secondary" onClick={this.handleCancel}>Cancel</Button>
              </CardFooter>
            </Card>
          </Col>
        </Row>
      </div>
    );
  }
}

export default RuleEditForm;
