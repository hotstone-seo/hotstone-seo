import React, { Component } from "react";

import {
  Button,
  Col,
  Form,
  FormGroup,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  Input,
  Label
} from "reactstrap";

class CanonicalForm extends Component {
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
    history.push("/base/MetatagPreview");
  }

  render() {
    const {
      visible,
      onCancel,
      onSave,
      canonical,
      action,
      onChange,
      languages,
      languageDefault
    } = this.props;
    let defaultValueLang =
      action !== "Add" ? canonical.locale : languageDefault;

    return (
      <Modal isOpen={visible}>
        <ModalHeader>{action} Canonical</ModalHeader>
        <ModalBody>
          <Form className="form-horizontal">
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Locale</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="select"
                  name="locale"
                  id="locale"
                  onChange={onChange.bind(this, "locale")}
                  defaultValue={defaultValueLang}
                  disabled
                >
                  {languages.map(ds => (
                    <option key={ds.lang_code} value={ds.lang_code}>
                      {ds.lang_code + "_" + ds.country_code}
                    </option>
                  ))}
                </Input>
              </Col>
            </FormGroup>
            <FormGroup row>
              {canonical !== undefined ? (
                <Input
                  type="hidden"
                  id="id"
                  name="id"
                  defaultValue={canonical.id}
                  onChange={onChange.bind(this, "id")}
                />
              ) : (
                ""
              )}
              <Col md="3">
                <Label htmlFor="text-input">Canonical Name</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="text"
                  id="canonical"
                  name="canonical"
                  placeholder="Canonical Name"
                  defaultValue={
                    canonical !== undefined ? canonical.canonical : ""
                  }
                  onChange={onChange.bind(this, "canonical")}
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

export default CanonicalForm;
