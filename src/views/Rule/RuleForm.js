import React from 'react';
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
} from 'reactstrap';

/* eslint-disable react/prefer-stateless-function */
class RuleForm extends React.Component {
  render() {
    const {
      visible, onCancel, onSave, rule
    } = this.props;

    return (
      <Modal
        isOpen={visible}
        title="Create a new Rule"
        okText="Save"
        onCancel={onCancel}
        onOk={onSave}
      >
        <Form className="form-horizontal"

        //props.addRule(rule)
        //setRule(initialFormState)
        >
          <FormGroup row>
            <Col md="3">
              <Label htmlFor="text-input">Name</Label>
            </Col>
            <Col xs="12" md="9">
              <Input type="hidden" id="id" name="id" />
              <Input type="text" id="name" name="name" placeholder="Name" value={rule.name} />
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
              <Label htmlFor="text-input">Data Source</Label>
            </Col>
            <Col xs="12" md="9">
              <Input type="select" name="datasource" id="datasource">
                <option value="1">Airport</option>
              </Input>
            </Col>
          </FormGroup>
        </Form>
      </Modal>
    );
  }
}
/* eslint-enable react/prefer-stateless-function */

export default RuleForm;