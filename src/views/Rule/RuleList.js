import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Table, Button, NavLink } from 'reactstrap';
import PropTypes from 'prop-types';
import axios from 'axios';

class RuleList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      rules: [],
    };

    this.columns = [
      {
        title: 'Rule ID',
        dataIndex: 'id',
        key: 'id',
      },
      {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
      },
      {
        title: 'URL Pattern',
        dataIndex: 'url_pattern',
        key: 'url_pattern',
      },
      {
        title: 'Updated At',
        dataIndex: 'updated_at',
        key: 'updated_at',
      },
      {
        title: 'Actions',
        key: 'actions',
        render: (text, record) => (
          <span>
            <Button type="link" onClick={() => this.handleClick(record)}>Edit</Button>
            <Divider type="vertical" />
            <Popconfirm title="Are you sure want to delete?" onConfirm={() => this.handleDelete(record.id)}>
              <Button type="link">Delete</Button>
            </Popconfirm>
          </span>
        ),
      },
    ];
    this.handleClick = this.handleClick.bind(this);
    this.handleEdit = this.handleEdit.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
  }

  componentDidMount() {
    axios.get('http://localhost:4000/rules')
      .then((res) => {
        const rules = res.data;
        this.setState({ rules });
      });
  }

  handleClick() {
    const { history } = this.props;
    history.push('/ruleForm');
  }
  
  handleEdit() {
    const { history } = this.props;
    history.push('/ruleEditForm');
  }

  handleDelete(id) {
    axios.delete(`http://localhost:4000/rules/${id}`)
      .then(() => {
        const { rules } = this.state;
        this.setState({ rules: rules.filter((env) => env.id !== id) });
      })
      .catch((error) => {
        message.error(error.message);
      });
  }

  render() {
    const { rules } = this.state;
   
    return (
      <div className="animated fadeIn">
        <Col xs="12" lg="12">
          <Card>
            <CardHeader>
              Rule
            </CardHeader>
            <CardBody>
              <div style={{ marginBottom: '.5rem' }}>
                <Button color="primary" onClick={this.handleClick}>Add New</Button>
              </div>
              <Table responsive bordered
                dataSource={rules}
                columns={this.columns}
                rowKey={(record) => record.id}
              />
              
              <Pagination>
                <PaginationItem><PaginationLink previous tag="button">Prev</PaginationLink></PaginationItem>
                <PaginationItem active>
                  <PaginationLink tag="button">1</PaginationLink>
                </PaginationItem>
                <PaginationItem><PaginationLink next tag="button">Next</PaginationLink></PaginationItem>
              </Pagination>
            </CardBody>
          </Card>
        </Col>
      </div>
    );
  }
}
RuleList.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};

export default RuleList;
