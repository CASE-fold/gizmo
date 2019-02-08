package service

import (
	"reflect"
	"testing"

	"github.com/CASE-fold/gizmo/v2/pubsub"
	"github.com/CASE-fold/gizmo/v2/pubsub/pubsubtest"
	"github.com/golang/protobuf/proto"

	"github.com/CASE-fold/gizmo/v2/examples/nyt"
)

func TestRun(t *testing.T) {
	tests := []struct {
		givenSub pubsub.Subscriber

		wantArticles []nyt.SemanticConceptArticle
	}{
		{
			&pubsubtest.TestSubscriber{
				ProtoMessages: []proto.Message{
					&nyt.SemanticConceptArticle{
						Url: "http://www.nytimes.com/awesome-cat-article-1",
					},
					&nyt.SemanticConceptArticle{
						Url: "http://www.nytimes.com/awesome-cat-article-2",
					},
				},
			},

			[]nyt.SemanticConceptArticle{
				nyt.SemanticConceptArticle{
					Url: "http://www.nytimes.com/awesome-cat-article-1",
				},
				nyt.SemanticConceptArticle{
					Url: "http://www.nytimes.com/awesome-cat-article-2",
				},
			},
		},
	}

	for _, test := range tests {

		// set test env
		sub = test.givenSub

		Run()

		for idx, article := range articles {
			if !reflect.DeepEqual(test.wantArticles[idx], article) {
				t.Errorf("got article[%d]:\n%#v\nexpected:\n%#v", idx, article, test.wantArticles[idx])
			}
		}

	}

}
