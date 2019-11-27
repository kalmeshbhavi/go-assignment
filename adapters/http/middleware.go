package http

import (
	"net/http"
)

func responseServerAdapter() ServerAdapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			resp := ToResponse(w)
			h.ServeHTTP(resp, r)
			//ctx := r.Context()
			err := resp.GetError()
			if err != nil {
				handleError(resp, r)
				return
			}
			actualResp := resp.GetResponse()
			if actualResp != nil {
				err := resp.WriteJSON(http.StatusOK, actualResp)
				if err != nil {

					/*logger.FromContext(ctx).Error("failed to write response", zap.Error(err))*/
				}
				return
			}
			//logger.FromContext(ctx).Error("nothing to write: both response and error are nil!")
		})
	}
}
