import React, { Component } from "react";

import {
  Button,
  Col,
  Form,
  FormGroup,
  Input,
  Label,
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter
} from "reactstrap";

class DataSourceForm extends Component {
  render() {
    const {
      visible,
      onCancel,
      onSave,
      datasource,
      action,
      onChange
    } = this.props;
    return (
      <Modal isOpen={visible}>
        <ModalHeader>{action} Data Source</ModalHeader>
        <ModalBody>
          <Form
            action=""
            method="post"
            encType="multipart/form-data"
            className="form-horizontal"
          >
            <FormGroup row>
              {datasource !== undefined ? (
                <Input
                  type="hidden"
                  id="id"
                  name="id"
                  defaultValue={datasource.id}
                  onChange={onChange.bind(this, "id")}
                />
              ) : (
                ""
              )}
              <Col md="3">
                <Label htmlFor="text-input">
                  Name <span className="mandatory-field">*</span>
                </Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="text"
                  id="name"
                  name="name"
                  placeholder="Name"
                  onChange={onChange.bind(this, "name")}
                  defaultValue={datasource !== undefined ? datasource.name : ""}
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="textarea-input">
                  URL <span className="mandatory-field">*</span>
                </Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="textarea"
                  name="url"
                  id="url"
                  rows="3"
                  placeholder="URL"
                  onChange={onChange.bind(this, "url")}
                  defaultValue={datasource !== undefined ? datasource.url : ""}
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Col xs="12" md="9">
                <span className="mandatory-field">* : must be filled in</span>
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
// DataSourceForm.propTypes = {
//   match: PropTypes.shape({
//     path: PropTypes.string,
//   }).isRequired,
// };
export default DataSourceForm;
