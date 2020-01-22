import React from "react";
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
} from "reactstrap";

class RuleForm extends React.Component {
  render() {
    const {
      visible,
      onCancel,
      onSave,
      rule,
      action,
      onChange,
      dataSources
    } = this.props;
    return (
      <Modal isOpen={visible}>
        <ModalHeader>{action} Rule</ModalHeader>
        <ModalBody>
          <Form className="form-horizontal">
            <FormGroup row>
              {rule !== undefined ? (
                <Input
                  type="hidden"
                  id="id"
                  name="id"
                  defaultValue={rule.id}
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
                  defaultValue={rule !== undefined ? rule.name : ""}
                  onChange={onChange.bind(this, "name")}
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
                  id="url_pattern"
                  name="url_pattern"
                  placeholder="URL Pattern"
                  defaultValue={rule !== undefined ? rule.url_pattern : ""}
                  onChange={onChange.bind(this, "url_pattern")}
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Data Source</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="select"
                  name="data_source_id"
                  id="data_source_id"
                  onChange={onChange.bind(this, "data_source_id")}
                  value={rule.data_source_id}
                >
                  <option value="-">-CHOOSE-</option>
                  {dataSources.map(ds => (
                    <option key={ds.id} value={ds.id}>
                      {ds.name}
                    </option>
                  ))}
                </Input>
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
export default RuleForm;
