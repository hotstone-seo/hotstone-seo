import React, { Component } from "react";
import PropTypes from "prop-types";

import {
  Button,
  Col,
  Form,
  FormGroup,
  Input,
  Label,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader
} from "reactstrap";

class MetatagForm extends Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.toggleFade = this.toggleFade.bind(this);
    this.state = {
      collapse: true,
      fadeIn: true,
      timeout: 300
    };

    this.handlePreview = this.handlePreview.bind(this);
  }

  toggle() {
    this.setState({ collapse: !this.state.collapse });
  }

  toggleFade() {
    this.setState(prevState => {
      return { fadeIn: !prevState };
    });
  }

  handlePreview() {
    const { history } = this.props;
    history.push("/metatagPreview");
  }

  render() {
    const { visible, onCancel, onSave, metatag, action, onChange } = this.props;
    return (
      <Modal isOpen={visible}>
        <ModalHeader>{action} Meta-Tag</ModalHeader>
        <ModalBody>
          <Input
            type="hidden"
            id="rule_id"
            name="rule_id"
            value="1"
            onChange={onChange.bind(this, "rule_id")}
          />
          <Form className="form-horizontal">
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Language</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="select"
                  name="language"
                  id="language"
                  onChange={onChange.bind(this, "locale_id")}
                >
                  <option value="1">ID</option>
                </Input>
              </Col>
            </FormGroup>
            <FormGroup row>
              {metatag !== undefined ? (
                <Input
                  type="hidden"
                  id="id"
                  name="id"
                  defaultValue={metatag.id}
                  onChange={onChange.bind(this, "id")}
                />
              ) : (
                ""
              )}
              <Col md="3">
                <Label htmlFor="text-input">Name</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="text"
                  id="name"
                  name="name"
                  placeholder="Name"
                  defaultValue={metatag !== undefined ? metatag.name : ""}
                  onChange={onChange.bind(this, "name")}
                />
              </Col>
            </FormGroup>

            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Content</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="text"
                  id="content"
                  name="content"
                  placeholder="content"
                  defaultValue={metatag !== undefined ? metatag.content : ""}
                  onChange={onChange.bind(this, "content")}
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Default Content</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="text"
                  id="default_content"
                  name="default_content"
                  placeholder="default content"
                  defaultValue={
                    metatag !== undefined ? metatag.default_content : ""
                  }
                />
              </Col>
            </FormGroup>
          </Form>
        </ModalBody>
        <ModalFooter>
          <Button color="warning" onClick={onSave}>
            Save
          </Button>{" "}
          <Button color="secondary" onClick={onCancel}>
            Cancel
          </Button>
        </ModalFooter>
      </Modal>
    );
  }
}
MetatagForm.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string
  }).isRequired
};
export default MetatagForm;
