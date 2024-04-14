package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestActor(t *testing.T) {
	t.Parallel()
	// Arrange a new actor for testing
	const (
		role1        = "role1"
		role2        = "role2"
		missingRole1 = "missingRole1"
	)
	actorID := uuid.New()
	roles := []string{role1, role2}
	actor := NewActor(actorID, roles)

	t.Run("HasOneOfRoles", func(t *testing.T) {
		t.Parallel()

		assert.True(t, actor.HasOneOfRoles(role1))
		assert.True(t, actor.HasOneOfRoles(role2))
		assert.True(t, actor.HasOneOfRoles(role1, missingRole1))
		assert.False(t, actor.HasOneOfRoles(missingRole1, missingRole1))
		assert.False(t, actor.HasOneOfRoles(missingRole1))

	})

	t.Run("HasOneOfRolesEmptyRoles", func(t *testing.T) {
		t.Parallel()

		// Test bad scenario
		assert.True(t, actor.HasOneOfRoles())
	})

	t.Run("HasOneOfRolesEmptyActor", func(t *testing.T) {
		t.Parallel()

		emptyActor := NewActor(actorID, make([]string, 0))

		assert.False(t, emptyActor.HasOneOfRoles(role1))
		assert.False(t, emptyActor.HasOneOfRoles(missingRole1))
	})

	t.Run("HasAllRoles", func(t *testing.T) {
		t.Parallel()

		assert.True(t, actor.HasAllRoles(role1, role2))
		assert.False(t, actor.HasAllRoles(role1, missingRole1))
		assert.False(t, actor.HasAllRoles(missingRole1))
	})

	t.Run("InitRoles", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, 2, actor.roles.Size(), "unexpected number of roles")
		assert.True(t, actor.roles.Contains(role1))
		assert.True(t, actor.roles.Contains(role2))
	})
}

func TestHasOneOfRoles_WithRole_ReturnsTrue(t *testing.T) {
	t.Parallel()

	actorID := uuid.New()
	roles := []string{"role1", "role2"}
	actor := NewActor(actorID, roles)

	assert.True(t, actor.HasOneOfRoles("role1"))
}

// Может так лучше?
func TestHasOneOfRoles_WithoutRole_ReturnsFalse(t *testing.T) {
	t.Parallel()

	actorID := uuid.New()
	roles := []string{"role1", "role2"}
	actor := NewActor(actorID, roles)

	assert.False(t, actor.HasOneOfRoles("role3"))
}

func TestHasOneOfRoles_WithEmptyRoles_ReturnsTrue(t *testing.T) {
	t.Parallel()

	actorID := uuid.New()
	roles := []string{"role1", "role2"}
	actor := NewActor(actorID, roles)

	assert.True(t, actor.HasOneOfRoles())
}

func TestHasOneOfRoles_WithNilRoles_ReturnsTrue(t *testing.T) {
	t.Parallel()

	actorID := uuid.New()
	roles := []string{"role1", "role2"}
	actor := NewActor(actorID, roles)

	assert.True(t, actor.HasOneOfRoles(nil...))
}

func TestGetRoles_ReturnsCorrectRoles(t *testing.T) {
	t.Parallel()

	actorID := uuid.New()
	roles := []string{"role1", "role2"}
	actor := NewActor(actorID, roles)

	returnedRoles := actor.GetRoles()

	assert.ElementsMatch(t, roles, returnedRoles)
}

func TestGetRoles_NoRoles_ReturnsEmptySlice(t *testing.T) {
	t.Parallel()

	actorID := uuid.New()
	actor := NewActor(actorID, nil)

	returnedRoles := actor.GetRoles()

	assert.Empty(t, returnedRoles)
}

func TestFillEmptyUUIDs_Course(t *testing.T) {
	t.Parallel()

	course := &Course{
		ID: uuid.Nil,
		Sections: []Section{
			{ID: uuid.Nil},
			{ID: uuid.New()},
		},
	}

	course.FillEmptyUUIDs()

	assert.NotEqual(t, uuid.Nil, course.ID, "Course ID should not be nil after FillEmptyUUIDs")
	assert.NotEqual(t, uuid.UUID{}, course.ID, "Course ID should not be empty after FillEmptyUUIDs")
	for _, section := range course.Sections {
		assert.NotEqual(t, uuid.Nil, section.ID, "Section ID should not be nil after FillEmptyUUIDs")
		assert.NotEqual(t, uuid.UUID{}, section.ID, "Section ID should not be empty after FillEmptyUUIDs")
	}
}

func TestFillEmptyUUIDs_Section(t *testing.T) {
	t.Parallel()

	section := &Section{
		ID: uuid.Nil,
		Themes: []Theme{
			{ID: uuid.Nil},
			{ID: uuid.New()},
		},
	}

	section.FillEmptyUUIDs()

	assert.NotEqual(t, uuid.Nil, section.ID, "Section ID should not be nil after FillEmptyUUIDs")
	assert.NotEqual(t, uuid.UUID{}, section.ID, "Section ID should not be empty after FillEmptyUUIDs")
	for _, theme := range section.Themes {
		assert.NotEqual(t, uuid.Nil, theme.ID, "Theme ID should not be nil after FillEmptyUUIDs")
		assert.NotEqual(t, uuid.UUID{}, theme.ID, "Theme ID should not be empty after FillEmptyUUIDs")
	}
}

func TestFillEmptyUUIDs_Theme(t *testing.T) {
	t.Parallel()

	theme := &Theme{
		ID: uuid.Nil,
	}

	theme.FillEmptyUUIDs()

	assert.NotEqual(t, uuid.Nil, theme.ID, "Theme ID should not be nil after FillEmptyUUIDs")
	assert.NotEqual(t, uuid.UUID{}, theme.ID, "Theme ID should not be empty after FillEmptyUUIDs")
}
