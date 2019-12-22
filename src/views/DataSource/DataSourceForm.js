import React, { Component } from 'react';
import PropTypes from 'prop-types';

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

class DataSourceForm extends Component {
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
    history.push('/dataSource');
  }

  render() {
    const {
      visible, onCancel, onSave, datasource, action, onChange
    } = this.props;
    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" md="9" lg="6">
            <Card>
              <CardHeader>
                <strong>Add New Data Source</strong>
              </CardHeader>
              <CardBody>
                <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Data Source Name</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="text" id="name" name="name" placeholder="Name" />
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Webhook</Label>
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
                            <DropdownItem>http://flight-service/airport</DropdownItem>
                          </DropdownMenu>
                        </InputGroupButtonDropdown>
                      </InputGroup>
                    </Col>
                  </FormGroup>

                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="textarea-input">Fields</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="textarea" name="fields" id="fields" rows="3"
                        placeholder="Fields" />
                    </Col>
                  </FormGroup>
                </Form>
              </CardBody>
              <CardFooter>
                <Button type="submit" size="md" color="primary" style={{ marginRight: "0.4em" }}><i className="fa fa-dot-circle-o"></i>Submit</Button>
                <Button type="button" size="md" color="secondary" onClick={this.handleCancel}>Cancel</Button>
              </CardFooter>
            </Card>
          </Col>
        </Row>
      </div>
    );
  }
}
DataSourceForm.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};
export default DataSourceForm;
