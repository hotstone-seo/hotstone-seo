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
  Input,
  Label,
  Row,
} from 'reactstrap';

class LanguageForm extends Component {
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
    history.push('/language');
  }

  render() {
    const {
      visible, onCancel, onSave, language, action, onChange
    } = this.props;

    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" md="9" lg="6">
            <Card>
              <CardHeader>
                <strong>Add New Language</strong>
              </CardHeader>
              <CardBody>
                <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Language Name</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="text" id="langName" name="langName" placeholder="Language Name" />
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Language Code</Label>
                    </Col>
                    <Col xs="12" md="9">
                        <Input type="text" id="langCode" name="langCode" placeholder="Language Code" />
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Country</Label>
                    </Col>
                    <Col xs="12" md="9">
                        <Input type="text" id="country" name="country" placeholder="Country" />
                    </Col>
                  </FormGroup>
                </Form>
              </CardBody>
              <CardFooter>
                <Button type="submit" size="md" color="primary" style={{ marginRight: "0.4em" }}><i className="fa fa-dot-circle-o"></i> Submit</Button>
                <Button type="button" size="md" color="secondary" onClick={this.handleCancel}>Cancel</Button>
              </CardFooter>
            </Card>
          </Col>
        </Row>
      </div>
    );
  }
}
LanguageForm.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};
export default LanguageForm;