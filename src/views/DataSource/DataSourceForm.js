import React, { Component } from 'react';
import PropTypes from 'prop-types';

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
  ModalFooter, 
} from 'reactstrap';

class DataSourceForm extends Component {
  render() {
    const {
      visible, onCancel, onSave, datasource, action, onChange
    } = this.props;
    return (
      <Modal isOpen={visible}>
        <ModalHeader>{action} Data Source</ModalHeader>
        <ModalBody>
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
                      <Label htmlFor="textarea-input">Fields</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="textarea" name="fields" id="fields" rows="3"
                        placeholder="Fields" />
                    </Col>
                  </FormGroup>
                </Form>
        </ModalBody>
        <ModalFooter>
          <Button color="warning" onClick={onSave}>Save</Button>{' '}
          <Button color="secondary" onClick={onCancel}>Cancel</Button>
        </ModalFooter>
      </Modal>
    );
  }
}
DataSourceForm.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};
export default DataSourceForm;
