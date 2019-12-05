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
  DropdownItem,
  DropdownMenu,
  DropdownToggle,
  Input,
  InputGroup,
  InputGroupButtonDropdown,
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
    history.push('/base/rule');
  }

  render() {
    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" md="9" lg="6">
            <Card>
              <CardHeader>
                <strong>Edit Rule</strong>
              </CardHeader>
              <CardBody>
                <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Name</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="text" id="name" name="name" placeholder="Name" />
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">URL Pattern</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="text" id="urlPattern" name="urlPattern" placeholder="URL Pattern" />
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="textarea-input">Exclusion</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="textarea" name="exclusion" id="exclusion" rows="3"
                        placeholder="Exclusion" />
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Data Source</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <InputGroup>
                        <InputGroupButtonDropdown addonType="prepend"
                          isOpen={this.state.first}
                          toggle={() => { this.setState({ first: !this.state.first }); }}>
                          <DropdownToggle caret color="primary">
                            -Choose-
                          </DropdownToggle>
                          <DropdownMenu className={this.state.first ? 'show' : ''}>
                            <DropdownItem>Airport</DropdownItem>
                          </DropdownMenu>
                        </InputGroupButtonDropdown>
                      </InputGroup>
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
