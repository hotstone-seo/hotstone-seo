import React, { Component } from 'react';

import {
  Form,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  Input,
  Col,
  FormGroup,
  Label,
  Button
} from 'reactstrap';

class LanguageForm extends Component {
  render() {
    const {
      visible, onCancel, onSave, language, action, onChange
    } = this.props;

    return (
      <Modal isOpen={visible}>
        <ModalHeader>{action} Language</ModalHeader>
        <ModalBody>
          <Form className="form-horizontal">
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Name</Label>
              </Col>
              <Col xs="12" md="9">
                {language !== undefined? (<Input type="hidden" id="id" name="id" defaultValue={language !== undefined?language.id:""} onChange={onChange.bind(this, 'id')}/>):""}
                <Input
                  type="text"
                  id="lang_code"
                  name="lang_code"
                  placeholder="Language Code"
                  defaultValue={language !== undefined?language.lang_code:""}
                  onChange={onChange.bind(this, 'lang_code')}
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Country Code</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="text"
                  id="country_code"
                  name="country_code"
                  placeholder="Country code"
                  defaultValue={language !== undefined?language.country_code:""}
                  onChange={onChange.bind(this, 'country_code')}
                />
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

export default LanguageForm;