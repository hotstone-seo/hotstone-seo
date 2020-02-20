import React, { Component } from "react";

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

class ScripttagForm extends Component {
  constructor(props) {
    super(props);
    this.handlePreview = this.handlePreview.bind(this);
  }

  handlePreview() {
    const { history } = this.props;
    history.push("/base/ScripttagPreview");
  }

  render() {
    const {
      visible,
      onCancel,
      onSave,
      scripttag,
      action,
      onChange,
      languageDefault,
      scriptTagPreviewValue,
      dataSources
    } = this.props;
    return (
      <Modal isOpen={visible}>
        <ModalHeader>{action} Script-Tag</ModalHeader>
        <ModalBody>
          <Form className="form-horizontal">
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Language</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="select"
                  name="locale"
                  id="locale"
                  onChange={onChange.bind(this, "locale")}
                  value={languageDefault}
                  disabled
                >
                  <option value="id_ID">id_ID</option>
                  <option value="en_US">en_US</option>
                </Input>
              </Col>
            </FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Type</Label>
              </Col>
              <Col xs="12" md="9">
                {scripttag !== undefined ? (
                  <Input
                    type="hidden"
                    id="id"
                    name="id"
                    defaultValue={scripttag !== undefined ? scripttag.id : ""}
                    onChange={onChange.bind(this, "id")}
                  />
                ) : (
                  ""
                )}
                <Input
                  type="select"
                  name="types"
                  id="types"
                  onChange={onChange.bind(this, "type")}
                  defaultValue={scripttag.type}
                >
                  <option value="-">-CHOOSE-</option>
                  <option value="Javascript">Javascript</option>
                </Input>
              </Col>
            </FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Source</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="text"
                  id="source"
                  name="source"
                  placeholder="Source"
                  defaultValue={scripttag !== undefined ? scripttag.source : ""}
                  onChange={onChange.bind(this, "source")}
                />
              </Col>
            </FormGroup>
            <FormGroup row></FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Preview</Label>
              </Col>
              <Col xs="12" md="9">
                <pre>{scriptTagPreviewValue}</pre>
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

export default ScripttagForm;
