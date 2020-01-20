import React from "react";
import { useForm } from "react-hook-form";
import { Machine, assign } from "xstate";
import { useMachine } from "@xstate/react";
import _ from "lodash";
import parse from "url-parse";

import useHotstoneAPI from "../../hooks/useHotstoneAPI";
import HotstoneAPI from "../../api/hotstone";
import inspectAxiosError, { isAxiosError } from "../../utils/axios";

const pageMachine = Machine({
  id: "simulation",
  initial: "idle",
  context: {
    url: null,
    matchResp: null,
    matchError: null
  },
  states: {
    idle: {},
    loading: {
      entry: assign({
        matchResp: null,
        matchError: null
      }),
      invoke: {
        src: (context, event) => HotstoneAPI.postProviderMatchRule(context.url),
        onDone: {
          target: "success",
          actions: assign({
            matchResp: (context, event) => {
              console.log("RESP: ", event);
              return event.data;
            }
          })
        },
        onError: {
          target: "failed",
          actions: assign({
            matchError: (context, event) => {
              console.log("ERR :", event);
              return event.data;
            }
          })
        }
      }
    },
    success: {},
    failed: {}
  },
  on: {
    SUBMIT: {
      target: ".loading",
      actions: assign({
        url: (context, event) => event.url
      })
    }
  }
});

function SimulationPage() {
  const [current, send] = useMachine(pageMachine);
  const { url, matchResp, matchError } = current.context;

  const { register, handleSubmit, errors } = useForm();
  const onSubmit = ({ url }) => {
    const urlObj = parse(url);

    send("SUBMIT", { url: urlObj.pathname });
  };

  return (
    <div className="container">
      <div className="row">
        <div className="col">
          <div className="card">
            <div className="card-header">Simulation: Rule Matching</div>
            <div className="card-body">
              <form
                onSubmit={handleSubmit(onSubmit)}
                className="form-horizontal needs-validation"
              >
                <div className="form-row">
                  <div className="col-md-9">
                    <input
                      name="url"
                      placeholder="URL"
                      className={"form-control " + (errors.url && "is-invalid")}
                      ref={register({
                        required: "Required"
                      })}
                    />
                    {errors.url && (
                      <div className="invalid-feedback">
                        {errors.url.message}
                      </div>
                    )}
                  </div>
                  <div className="col-auto">
                    <button
                      type="submit"
                      className="btn btn-primary"
                      disabled={current.matches("loading")}
                    >
                      Submit
                    </button>
                  </div>
                </div>
              </form>
            </div>

            {current.matches("loading") && (
              <div className="card-footer">
                <div className="alert alert-warning" role="alert">
                  Loading ...
                </div>
              </div>
            )}

            {renderIfSuccess(matchResp)}
            {renderIfError(matchError)}
          </div>
        </div>
      </div>
      <div className="row">
        <div className="col"></div>
        <div className="col"></div>
        <div className="col"></div>
      </div>
      <div className="row">
        <div className="col"></div>
      </div>
    </div>
  );
}

function renderIfSuccess(matchResp) {
  if (matchResp) {
    const { rule_id, path_param } = matchResp.data;
    return (
      <div className="card-footer">
        <div className="alert alert-success" role="alert">
          Matched. <br />
          {!_.isEmpty(path_param) && (
            <div>
              <div>Path params:</div>
              <ul>
                {Object.entries(path_param).map(([key, value]) => {
                  return (
                    <li key={key}>
                      {key}: {value}
                    </li>
                  );
                })}
              </ul>
            </div>
          )}
        </div>
      </div>
    );
  }
}

function renderIfError(matchError) {
  if (matchError) {
    let msgError = matchError.message;
    if (matchError.response) {
      msgError = matchError.response.data.message;
    }

    return (
      <div className="card-footer">
        <div className="alert alert-danger" role="alert">
          {msgError}
        </div>
      </div>
    );
  }
}

export default SimulationPage;
