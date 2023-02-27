package playlist

import (
	"github.com/stretchr/testify/require"
	"gocloudcamp/internal/domain/song"
	"testing"
	"time"
)

type ferr func() error

func Test_PlaylistEmpty_PlayPauseNextPrev_Fail(t *testing.T) {
	pl := NewPlaylist()
	require.True(t, pl.Empty())

	for _, tc := range []struct {
		name string
		f    ferr
	}{
		{"play", pl.Play},
		{"pause", pl.Pause},
		{"next", pl.Next},
		{"prev", pl.Prev},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Error(t, tc.f())
			require.False(t, pl.Played())
		})
	}
}

func Test_PlaylistOneSong_PlayPauseNextPrev_Success(t *testing.T) {
	pl := NewPlaylist()
	expected := song.Song{Name: "first", Duration: time.Second}
	pl.AddSong(&expected)
	require.False(t, pl.Empty())

	for _, tc := range []struct {
		name   string
		chain  []ferr
		played bool
	}{
		{
			"play_pause",
			[]ferr{pl.Play, pl.Pause},
			false,
		},
		{
			"play_pause_next",
			[]ferr{pl.Play, pl.Pause, pl.Next},
			false,
		},
		{
			"play_pause_prev",
			[]ferr{pl.Play, pl.Pause, pl.Prev},
			false,
		},
		{
			"play_pause_next_prev_next_prev",
			[]ferr{pl.Play, pl.Pause, pl.Next, pl.Prev, pl.Next, pl.Prev},
			false,
		},
		{
			"play_next",
			[]ferr{pl.Play, pl.Next},
			true,
		},
		{
			"play_prev",
			[]ferr{pl.Play, pl.Prev},
			true,
		},
		{
			"play_next_prev_next_prev",
			[]ferr{pl.Play, pl.Next, pl.Prev, pl.Next, pl.Prev},
			true,
		},
	} {
		_ = pl.Pause()
		_ = pl.First()

		t.Run(tc.name, func(t *testing.T) {
			for _, f := range tc.chain {
				require.NoError(t, f())
			}

			require.Equal(t, tc.played, pl.Played())

			actual, errCurrSong := pl.CurrentSong()
			require.NoError(t, errCurrSong)
			require.Equal(t, expected.Name, actual.Name)
		})
	}
}

func Test_PlaylistSomeSongs_PlayPauseNextPrev_Success(t *testing.T) {
	pl := NewPlaylist()
	expected := []song.Song{
		{Name: "first", Duration: time.Second},
		{Name: "second", Duration: time.Second},
		{Name: "third", Duration: time.Second},
	}
	for idx := range expected {
		pl.AddSong(&expected[idx])
	}
	require.False(t, pl.Empty())

	for _, tc := range []struct {
		name     string
		chain    []ferr
		played   bool
		expected string
	}{
		{
			"play_next_pause",
			[]ferr{pl.Play, pl.Next, pl.Pause},
			false,
			"second",
		},
		{
			"play_prev_pause",
			[]ferr{pl.Play, pl.Prev, pl.Pause},
			false,
			"third",
		},
		{
			"play_next_next_next",
			[]ferr{pl.Play, pl.Next, pl.Next, pl.Next},
			true,
			"first",
		},
	} {
		_ = pl.Pause()
		_ = pl.First()

		t.Run(tc.name, func(t *testing.T) {
			for _, f := range tc.chain {
				require.NoError(t, f())
			}

			require.Equal(t, tc.played, pl.Played())

			actual, errCurrSong := pl.CurrentSong()
			require.NoError(t, errCurrSong)
			require.Equal(t, tc.expected, actual.Name)
		})
	}
}
