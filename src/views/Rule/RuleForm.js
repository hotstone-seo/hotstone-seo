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
  Button
} from 'reactstrap';

class RuleForm extends React.Component {
  render() {
    const {
      visible, onCancel, onSave, rule, action, onChange
    } = this.props;
     
    return (
      <Modal isOpen={visible}>
        <ModalHeader>{action} Rule</ModalHeader>
        <ModalBody>
          <Form className="form-horizontal">
 
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Name</Label>
              </Col>
              <Col xs="12" md="9">
                {rule !== undefined? (<Input type="hidden" id="id" name="id" />):""}
                
                <Input
                  type="text"
                  id="name"
                  name="name"
                  placeholder="Name"
                  defaultValue={rule !== undefined?rule.name:""}
                  onChange={onChange.bind(this, 'name')}
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">URL Pattern</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="text"
                  id="urlPattern"
                  name="urlPattern"
                  placeholder="URL Pattern"
                  defaultValue={rule !== undefined?rule.url_pattern:""}
                  onChange={onChange.bind(this, 'urlPattern')}
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Data Source</Label>
              </Col>
              <Col xs="12" md="9">
                <Input type="select" name="datasource" id="datasource" onChange={onChange.bind(this, 'datasource')}>
                  <option value="1" selected>Airport</option>
                </Input>
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
//export default Form.create()(RuleForm)
export default RuleForm;