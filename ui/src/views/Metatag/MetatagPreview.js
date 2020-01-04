import React, { Component } from 'react';
import PropTypes from 'prop-types';

import {
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Col,
  FormGroup,
  Row,
} from 'reactstrap';
var thisIsMyCopy = '<head><meta charset="UTF-8"><meta name="description" content="Free Web tutorials"></head>';

class MetatagPreview extends Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.toggleFade = this.toggleFade.bind(this);
    this.state = {
      collapse: true,
      fadeIn: true,
      timeout: 300
    };
    this.handleBack = this.handleBack.bind(this);
  }

  toggle() {
    this.setState({ collapse: !this.state.collapse });
  }

  toggleFade() {
    this.setState((prevState) => { return { fadeIn: !prevState } });
  }
  handleBack() {
    const { history } = this.props;
    history.push('/metatagForm');
  }
  render() {
    return (
      <div className="animated fadeIn">

        <Row>
           <Col xs="12" md="9" lg="6">
            <Card>
              <CardHeader>
                <strong> Meta-Tag Preview</strong>
              </CardHeader>
              <CardBody>

                <FormGroup row>
                  <Col xs="12" md="12">

                    <div className="content" > {thisIsMyCopy}</div>
                  </Col>
                </FormGroup>

              </CardBody>
              <CardFooter>
                <Button type="button" size="sm" color="secondary" onClick={this.handleBack}><i className="fa fa-ban"></i> Back</Button>
              </CardFooter>
            </Card>
          </Col>
        </Row>
      </div>
    );
  }
}

MetatagPreview.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};

export default MetatagPreview;