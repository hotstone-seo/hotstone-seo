import React from "react";
import { useForm } from "react-hook-form";
import { Machine, assign } from "xstate";
import { useMachine } from "@xstate/react";
import { HotStoneClient } from "hotstone-client";
import { Link } from "react-router-dom";
import _ from "lodash";
import parse from "url-parse";

import HotstoneAPI from "../../api/hotstone";

const client = new HotStoneClient(process.env.REACT_APP_API_URL);

const pageMachine = Machine({
  id: "simulation",
  initial: "init",
  context: {
    url: null,
    locale: "en-US",
    listLocale: [],
    matchResp: null,
    matchError: null,
    pageError: null
  },
  states: {
    idle: {},
    init: {
      invoke: {
        src: context => HotstoneAPI.getLocales(),
        onDone: {
          target: "idle",
          actions: assign({
            listLocale: (context, event) => {
              console.log("[init] RESP: ", event);
              const listLocale = event.data.data;
              return listLocale.map(({ lang_code, country_code }) => {
                return `${lang_code}-${country_code}`;
              });
            }
          })
        },
        onError: {
          target: "failed",
          actions: assign({
            pageError: (context, event) => {
              console.log("[init] ERR :", event);
              return event.data;
            }
          })
        }
      }
    },
    submitting: {
      entry: assign({
        matchResp: null,
        matchError: null
      }),
      invoke: {
        src: context => matchThenGetTags(client, context.locale, context.url),
        onDone: {
          target: "success",
          actions: assign({
            matchResp: (context, event) => {
              console.log("RESP: ", event);
              const rule = event.data;
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
      target: ".submitting",
      actions: assign({
        url: (context, event) => event.url,
        locale: (context, event) => event.locale
      })
    }
  }
});

async function matchThenGetTags(client, locale, url) {
  const rule = await client.match(url);
  if (_.isEmpty(rule)) {
    throw new Error("Not matched");
  }
  const tags = await client.tags(rule, locale);
  const data = { rule, tags };
  return data;
}

function SimulationPage() {
  const [current, send] = useMachine(pageMachine);
  const { matchResp, matchError } = current.context;

  const { register, handleSubmit, errors } = useForm();
  const onSubmit = ({ locale, url }) => {
    const urlObj = parse(url);

    send("SUBMIT", { locale, url: urlObj.pathname });
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

                  {!_.isEmpty(current.context.listLocale) && (
                    <div className="col-auto">
                      <select
                        name="locale"
                        className="form-control"
                        defaultValue={current.context.locale}
                        ref={register({ required: true })}
                      >
                        {current.context.listLocale.map((locale, index) => {
                          return (
                            <option key={index} value={locale}>
                              {locale}
                            </option>
                          );
                        })}
                      </select>
                    </div>
                  )}

                  <div className="col-auto">
                    <button
                      type="submit"
                      className="btn btn-primary"
                      disabled={isLoading(current)}
                    >
                      Submit
                    </button>
                  </div>
                </div>
              </form>
            </div>

            {isLoading(current) && (
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
    const { rule_id, path_param } = matchResp.rule;
    const tags = matchResp.tags;

    let rawHtmlPreview;
    if (_.isEmpty(tags)) {
      rawHtmlPreview = (
        <div>
          No tags data. Register tags at{" "}
          <Link to={`/rule-detail/?id=${rule_id}`}>Rule Detail</Link>
        </div>
      );
    } else {
      rawHtmlPreview = <pre>TODO</pre>;
    }

    return (
      <div className="card-body">
        <div className="alert alert-success" role="alert">
          Matched (<Link to={`/rule-detail/?id=${rule_id}`}>Rule Detail</Link>)
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
        <strong>Raw HTML Tags Preview</strong>
        {rawHtmlPreview}
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

function isLoading(current) {
  return current.matches("init") || current.matches("submitting");
}

export default SimulationPage;
