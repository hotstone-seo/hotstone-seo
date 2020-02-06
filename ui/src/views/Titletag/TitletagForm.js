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

class TitletagForm extends Component {
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
    history.push("/titletagPreview");
  }

  render() {
    const {
      visible,
      onCancel,
      onSave,
      titletag,
      action,
      onChange,
      languages,
      languageDefault,
      titleTagPreviewValue
    } = this.props;
    var defaultValueLang = action !== "Add" ? titletag.locale : languageDefault;
    //if (defaultValueLang !== null)
    //  defaultValueLang = defaultValueLang.toLowerCase();
    return (
      <Modal isOpen={visible}>
        <ModalHeader>{action} Title-Tag</ModalHeader>
        <ModalBody>
          <Form className="form-horizontal">
            {titletag !== undefined ? (
              <Input
                type="hidden"
                id="id"
                name="id"
                defaultValue={titletag !== undefined ? titletag.id : ""}
                onChange={onChange.bind(this, "id")}
              />
            ) : (
              ""
            )}
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
                    <option
                      key={ds.lang_code}
                      value={ds.lang_code + "-" + ds.country_code}
                    >
                      {ds.lang_code + "-" + ds.country_code}
                    </option>
                  ))}
                </Input>
              </Col>
            </FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Title</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="text"
                  id="title"
                  name="title"
                  placeholder="title"
                  defaultValue={titletag !== undefined ? titletag.name : ""}
                  onChange={onChange.bind(this, "title")}
                />
              </Col>
            </FormGroup>
            <FormGroup row></FormGroup>
            <FormGroup row>
              <Col md="3">
                <Label htmlFor="text-input">Preview</Label>
              </Col>
              <Col xs="12" md="9">
                <pre>{titleTagPreviewValue}</pre>
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

export default TitletagForm;
