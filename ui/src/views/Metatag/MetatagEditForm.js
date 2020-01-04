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

class MetatagEditForm extends Component {
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
    history.push('/metatag');
  }

  render() {
    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" md="9" lg="6">
            <Card>
              <CardHeader>
                <strong>Edit Meta-Tag</strong>
              </CardHeader>
              <CardBody>
                <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">

                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Name</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="text" id="name" name="name" placeholder="Name" value="" />

                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Content</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="text" id="content" name="content" placeholder="Content" value="" />

                    </Col>
                  </FormGroup>

                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Default Content</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="text" id="defaultcontent" name="defaultcontent" placeholder="Default Content" />
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Rule</Label>
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
                            <DropdownItem>Airport Detail</DropdownItem>
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
MetatagEditForm.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};
export default MetatagEditForm;